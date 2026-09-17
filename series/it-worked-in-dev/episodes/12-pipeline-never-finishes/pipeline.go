// Package pipelineneverfinishes compares three ways of answering the question
// a pipeline runner has to answer before it starts: does following "what runs
// next" from here ever come back to something it has already run?
package pipelineneverfinishes

// Stage is one step of a pipeline. Each stage names the single stage that runs
// after it, and the last one names nothing.
//
// A redirect chain, a retry escalation, a symlink, and a workflow step all have
// this shape: one forward pointer, and a config file where somebody can point
// the last one back at the middle.
type Stage struct {
	Name string
	Next *Stage
}

// HasCycleBySeen remembers every stage it has walked past. If it arrives
// somewhere it has already been, the pipeline loops.
//
// This is the definition of "loops" written down: been here before. It needs no
// argument about why it works, and it is the version that can also say where
// the loop starts, which is what an operator actually wants.
func HasCycleBySeen(head *Stage) bool {
	seen := map[*Stage]bool{}
	for s := head; s != nil; s = s.Next {
		if seen[s] {
			return true
		}
		seen[s] = true
	}
	return false
}

// EntryBySeen returns the name of the first stage that repeats - the top of the
// loop - or "" when the pipeline terminates.
func EntryBySeen(head *Stage) string {
	seen := map[*Stage]bool{}
	for s := head; s != nil; s = s.Next {
		if seen[s] {
			return s.Name
		}
		seen[s] = true
	}
	return ""
}

// HasCycleByHopLimit gives up after limit steps and calls that a loop. It is
// what most HTTP clients do with redirects, and it is wrong in a way that only
// shows up on somebody else's data: a legitimate chain longer than the limit is
// reported as a cycle.
//
// TestHopLimitLiesAboutLongChains holds the pipeline it is wrong on.
func HasCycleByHopLimit(head *Stage, limit int) bool {
	hops := 0
	for s := head; s != nil; s = s.Next {
		if hops >= limit {
			return true
		}
		hops++
	}
	return false
}

// HasCycleByTwoPointers walks two cursors, one stage at a time and two stages
// at a time. If the chain ends, the fast one runs off it. If it loops, the fast
// one is inside the loop with the slow one and closes the gap by exactly one
// stage per turn, so they land on the same stage rather than stepping over each
// other.
func HasCycleByTwoPointers(head *Stage) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // the same stage, not the same name
			return true
		}
	}
	return false // the chain has an end, so nothing points back
}

// EntryByTwoPointers finds where the loop starts without remembering anything:
// once the two cursors meet, a third walk from the head moves in step with one
// of them and they meet at the top of the loop.
//
// It costs a second pass over the chain, which is the price of not having the
// map that already knew.
func EntryByTwoPointers(head *Stage) string {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// The distance from the head to the loop's top equals the
			// distance from the meeting point to the loop's top, going
			// forward. So walk both, one stage at a time.
			at := head
			for at != slow {
				at = at.Next
				slow = slow.Next
			}
			return at.Name
		}
	}
	return ""
}

// Length counts the stages a terminating pipeline holds, for the benchmark to
// report the shape it is running on. It never returns on a looping pipeline,
// which is the whole problem in one sentence.
func Length(head *Stage) int {
	n := 0
	for s := head; s != nil; s = s.Next {
		n++
	}
	return n
}
