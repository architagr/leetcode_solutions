package runningminwithoutrescan

// A budget plan: an opening balance and a list of planned transactions, in
// order. After every change the planner shows the lowest balance the plan
// ever reaches - the overdraft warning. Changes are "add a transaction at the
// end" and "undo the last one".

// PlanByRescan is what I would write. Keep the transactions; to answer, walk
// them from the opening balance and track the lowest point.
type PlanByRescan struct {
	Opening int
	tx      []int
}

func (p *PlanByRescan) Add(amount int) { p.tx = append(p.tx, amount) }

func (p *PlanByRescan) Undo() {
	if len(p.tx) > 0 {
		p.tx = p.tx[:len(p.tx)-1]
	}
}

func (p *PlanByRescan) Lowest() int {
	bal, low := p.Opening, p.Opening
	for _, a := range p.tx {
		bal += a
		if bal < low {
			low = bal
		}
	}
	return low
}

// PlanByOneMin is the fix that looks obvious: keep the balance and the
// lowest point as fields, update them on Add. Undo subtracts the amount
// back out of the balance - and has no way to restore the lowest point if
// the transaction being undone was the one that set it.
type PlanByOneMin struct {
	Opening  int
	tx       []int
	bal, low int
	started  bool
}

func (p *PlanByOneMin) init() {
	if !p.started {
		p.bal, p.low, p.started = p.Opening, p.Opening, true
	}
}

func (p *PlanByOneMin) Add(amount int) {
	p.init()
	p.tx = append(p.tx, amount)
	p.bal += amount
	if p.bal < p.low {
		p.low = p.bal
	}
}

func (p *PlanByOneMin) Undo() {
	p.init()
	if len(p.tx) > 0 {
		p.bal -= p.tx[len(p.tx)-1]
		p.tx = p.tx[:len(p.tx)-1]
		// p.low stays where it was: the plan may no longer go that low.
	}
}

func (p *PlanByOneMin) Lowest() int { p.init(); return p.low }

// entry is one transaction and the plan as it stood right after it: the
// balance, and the lowest the balance had been up to and including it.
type entry struct {
	amount, bal, low int
}

// PlanByEntry keeps, per transaction, the answer at that point. Undo drops
// an entry and the answer below it is already there.
type PlanByEntry struct {
	Opening int
	e       []entry
}

func (p *PlanByEntry) Add(amount int) {
	bal, low := p.Opening, p.Opening
	if n := len(p.e); n > 0 {
		bal, low = p.e[n-1].bal, p.e[n-1].low
	}
	bal += amount
	low = min(low, bal) // the lowest so far is the lowest before, or now
	p.e = append(p.e, entry{amount, bal, low})
}

func (p *PlanByEntry) Undo() {
	if len(p.e) > 0 {
		p.e = p.e[:len(p.e)-1]
	}
}

func (p *PlanByEntry) Lowest() int {
	if n := len(p.e); n > 0 {
		return p.e[n-1].low
	}
	return p.Opening
}
