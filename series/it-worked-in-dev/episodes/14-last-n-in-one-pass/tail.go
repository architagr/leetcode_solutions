package lastninonepass

import "iter"

// A crash reporter attaches the last N lines of the request log to the report.
//
// The log is a stream: a bufio.Scanner over a file, a cursor over a table, a
// channel fed by a tailer. It yields lines in order and it does not know how
// many there are, which is the whole difficulty - "the last 200" is defined
// relative to an end nobody has reached yet.
type Log = iter.Seq[string]

// LastNByCollecting is what I have written, twice, in two languages.
//
// Read it all, then take the tail. It is one idea, it reads exactly like the
// sentence in the ticket, and it is correct for every input including a log
// shorter than n.
func LastNByCollecting(log Log, n int) []string {
	var all []string
	for line := range log {
		all = append(all, line)
	}
	if len(all) <= n {
		return all
	}
	return all[len(all)-n:]
}

// LastNByCollectingAndCopying is the same function with the one-line fix for
// the part of it that surprises people: a slice of a slice shares the backing
// array, so the tail returned above keeps every line of the log reachable
// long after the function returns.
func LastNByCollectingAndCopying(log Log, n int) []string {
	tail := LastNByCollecting(log, n)
	out := make([]string, len(tail))
	copy(out, tail)
	return out
}

// LastNByCountingTwice is the version you write when somebody points out the
// memory. Walk it once to learn the length, then walk it again and keep from
// index count-n.
//
// It needs `open` rather than a Log, because walking twice means starting
// twice. A file allows that. A socket, a channel from a tailer, and a
// paginated API with a one-shot cursor do not, and neither does a log that is
// still being written to - the second walk sees a different log.
func LastNByCountingTwice(open func() Log, n int) []string {
	count := 0
	for range open() {
		count++
	}
	from := count - n
	if from < 0 {
		from = 0
	}
	out := make([]string, 0, min(n, count))
	i := 0
	for line := range open() {
		if i >= from {
			out = append(out, line)
		}
		i++
	}
	return out
}

// LastNByRing keeps exactly n lines and overwrites the oldest.
//
// w is where the next line goes, which is also where the oldest line currently
// is - the read cursor trails the write cursor by exactly n, and that fixed
// gap is the length the function never computes.
func LastNByRing(log Log, n int) []string {
	if n <= 0 {
		return nil
	}
	buf := make([]string, 0, n)
	w, count := 0, 0
	for line := range log {
		if len(buf) < n {
			buf = append(buf, line)
		} else {
			buf[w] = line // the line this evicts is exactly n lines old
		}
		w = (w + 1) % n
		count++
	}
	if count <= n {
		return buf
	}
	// Unwrap: the oldest line is at w, so the answer is buf[w:] then buf[:w].
	out := make([]string, 0, n)
	out = append(out, buf[w:]...)
	out = append(out, buf[:w]...)
	return out
}
