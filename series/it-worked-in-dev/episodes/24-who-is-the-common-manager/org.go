package whoisthecommonmanager

// The org chart as the HR service returns it: every employee with their
// direct reports. There is no manager pointer.
type Employee struct {
	ID      int
	Reports []*Employee

	// In and Out are filled by Number: an employee's whole reporting line
	// sits between their In and Out. Zero until Number runs.
	In, Out int
}

// PathTo is the function the directory page already had, behind the
// "reports to" breadcrumb: search from the top and return the chain of
// managers down to the employee, or nil if they are not in this org.
func PathTo(root *Employee, id int) []*Employee {
	if root == nil {
		return nil
	}
	if root.ID == id {
		return []*Employee{root}
	}
	for _, r := range root.Reports {
		if p := PathTo(r, id); p != nil {
			return append([]*Employee{root}, p...)
		}
	}
	return nil
}

// ManagerByPaths is what I would write. Everyone's chain of managers starts
// at the CEO; the lowest common manager is the last person on all of them.
//
// One breadcrumb per attendee, then the longest shared prefix.
func ManagerByPaths(root *Employee, ids []int) *Employee {
	var common []*Employee
	for i, id := range ids {
		p := PathTo(root, id)
		if p == nil {
			return nil
		}
		if i == 0 {
			common = p
			continue
		}
		k := 0
		for k < len(common) && k < len(p) && common[k] == p[k] {
			k++
		}
		common = common[:k]
	}
	if len(common) == 0 {
		return nil
	}
	return common[len(common)-1]
}

// ManagerByCount walks the chart once, bottom-up. Every employee reports how
// many of the attendees are in their reporting line, themselves included; the
// first one to reach all of them is the lowest common manager.
func ManagerByCount(root *Employee, ids []int) *Employee {
	want := map[int]bool{}
	for _, id := range ids {
		want[id] = true
	}
	need := len(want)
	var answer *Employee
	var walk func(e *Employee) int
	walk = func(e *Employee) int {
		n := 0
		if want[e.ID] {
			n++ // an employee is in their own reporting line
		}
		for _, r := range e.Reports {
			n += walk(r)
			if answer != nil {
				return n // found below: stop, or every manager above overwrites it
			}
		}
		if n == need {
			answer = e // bottom-up, so the first to see everyone is the lowest
		}
		return n
	}
	if need > 0 && root != nil {
		walk(root)
	}
	return answer
}

// Number gives every employee a pre-order In and an Out, so that someone's
// whole reporting line is exactly the employees numbered In..Out. One walk,
// done when the chart changes rather than per question.
func Number(root *Employee) map[int]*Employee {
	byID := map[int]*Employee{}
	next := 0
	var walk func(e *Employee)
	walk = func(e *Employee) {
		byID[e.ID] = e
		e.In = next
		next++
		for _, r := range e.Reports {
			walk(r)
		}
		e.Out = next - 1
	}
	walk(root)
	return byID
}

// ManagerByNumbers answers from the numbering: the attendees span In values
// from lo to hi, and the lowest common manager is the deepest employee whose
// range covers [lo, hi]. Walk down from the top into whichever report covers
// it - the same move as a BST, with a range in place of a value.
func ManagerByNumbers(root *Employee, byID map[int]*Employee, ids []int) *Employee {
	if len(ids) == 0 {
		return nil
	}
	lo, hi := -1, -1
	for _, id := range ids {
		e := byID[id]
		if e == nil {
			return nil
		}
		if lo < 0 || e.In < lo {
			lo = e.In
		}
		if e.In > hi {
			hi = e.In
		}
	}
	e := root
	for {
		down := (*Employee)(nil)
		for _, r := range e.Reports {
			if r.In <= lo && hi <= r.Out {
				down = r
				break
			}
		}
		if down == nil {
			return e // no single report covers everyone: they meet here
		}
		e = down
	}
}
