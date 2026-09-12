// Package pricingtierlookup compares two ways of finding which band a usage
// figure falls into, when the bands are already sorted.
package pricingtierlookup

// Tier is one band of a usage-based price list: everything from MinUnits up to
// the next tier's MinUnits is charged at this rate.
type Tier struct {
	MinUnits  int
	PricePerK int // in cents, per thousand units
	Name      string
}

// TierForScan returns the tier a usage figure falls into by walking the list.
//
// This is the version that gets written first, and it is the clearest possible
// statement of what a tier is: keep taking tiers while they still apply, and
// the last one that did is the answer. It makes no assumption about the list
// beyond it being in order, and it reads the same way the pricing page does.
func TierForScan(tiers []Tier, units int) (Tier, bool) {
	var found Tier
	var ok bool
	for _, t := range tiers {
		if t.MinUnits > units {
			break
		}
		found, ok = t, true
	}
	return found, ok
}

// TierForBinary answers the same question by halving the range instead of
// walking it.
//
// The list is sorted by MinUnits, which means one comparison rules out
// everything on one side of it. lo ends up on the last tier whose MinUnits is
// at or below units.
func TierForBinary(tiers []Tier, units int) (Tier, bool) {
	if len(tiers) == 0 || tiers[0].MinUnits > units {
		return Tier{}, false
	}
	lo, hi := 0, len(tiers)-1
	for lo < hi {
		// Bias the midpoint upward. With lo and hi adjacent, rounding down
		// would pick lo, and the lo = mid branch would never advance - the
		// loop would spin forever on two remaining candidates.
		mid := (lo + hi + 1) / 2
		if tiers[mid].MinUnits <= units {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return tiers[lo], true
}
