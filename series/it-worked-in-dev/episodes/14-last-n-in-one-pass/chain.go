package lastninonepass

// The same question on data that is already in memory, which is where the gap
// costs nothing at all to hold.
//
// Line is one entry of an in-memory ring of log records - the kind a server
// keeps so a crash report has something to attach. It is a chain: each entry
// names the next one, and nothing can walk backwards.
type Line struct {
	Text string
	Next *Line
}

// NthFromEndByCounting is the version day 24 of the challenge publishes, and
// the version almost everyone writes: count the chain, then walk to
// count-n from the front.
func NthFromEndByCounting(head *Line, n int) *Line {
	count := 0
	for l := head; l != nil; l = l.Next {
		count++
	}
	if n <= 0 || n > count {
		return nil
	}
	l := head
	for i := 0; i < count-n; i++ {
		l = l.Next
	}
	return l
}

// NthFromEndByGap never learns the length.
//
// Move `lead` n entries forward, then move both until lead falls off the end.
// The gap between them never changes, so when lead is at the end, behind is n
// from it. Two variables, one pass, and nothing held.
func NthFromEndByGap(head *Line, n int) *Line {
	if n <= 0 {
		return nil
	}
	lead := head
	for i := 0; i < n; i++ {
		if lead == nil {
			return nil // the chain is shorter than n
		}
		lead = lead.Next
	}
	behind := head
	for lead != nil {
		lead = lead.Next
		behind = behind.Next
	}
	return behind
}
