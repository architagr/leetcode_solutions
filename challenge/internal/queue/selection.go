// Selection picks which queued entry goes out next to a given
// destination, and guards against a scheduled run posting twice in one
// day. It lives here rather than in a per-network package because every
// destination selects the same way — the network only differs in how the
// content is delivered.

package queue

import "time"

// SelectNext returns the oldest entry, by day, that has not yet been
// posted to destination. ok is false when there is nothing left to post,
// which is a normal idle day rather than an error.
func SelectNext(q *Queue, destination string) (entry Entry, ok bool) {
	for _, e := range q.Entries {
		if e.IsPosted(destination) {
			continue
		}
		// Entries are ordered by day in practice, but selection must not
		// depend on file order — a hand-edited queue shouldn't change
		// which day goes out next.
		if !ok || e.Day < entry.Day {
			entry, ok = e, true
		}
	}
	return entry, ok
}

// PostedOn returns the entry, if any, that already went out to
// destination on the same UTC calendar day as now.
//
// A scheduled run can fire more than once a day — GitHub reruns a
// delayed cron, a dispatch retries a run that already succeeded — and
// the queue alone can't tell those apart from a legitimate next day,
// since it only records that an entry was posted, not when it was due.
// Comparing against the wall clock is what makes a second run that day
// a no-op instead of burning tomorrow's entry.
//
// Timestamps are compared in UTC because that's how MarkPosted writes
// them; an entry whose stamp doesn't parse is treated as not-today
// rather than blocking the day's post on a malformed field.
func PostedOn(q *Queue, destination string, now time.Time) (entry Entry, ok bool) {
	y, m, d := now.UTC().Date()
	for _, e := range q.Entries {
		if !e.IsPosted(destination) {
			continue
		}
		at, err := time.Parse(time.RFC3339, *e.PostedAt[destination])
		if err != nil {
			continue
		}
		ay, am, ad := at.UTC().Date()
		if ay != y || am != m || ad != d {
			continue
		}
		// Report the newest day sent today, so a message naming it reads
		// as the post that actually just went out.
		if !ok || e.Day > entry.Day {
			entry, ok = e, true
		}
	}
	return entry, ok
}
