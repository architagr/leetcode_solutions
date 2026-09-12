// Package dedupeuserlist compares two ways of finding which records in one
// list are missing from another - the nested loop everyone writes first, and
// the same job with a set.
package dedupeuserlist

// User is deliberately a struct rather than a bare string. Real de-duplication
// runs over records, and the field being matched on is usually not the whole
// thing, which is what makes the naive version look reasonable.
type User struct {
	ID    string
	Email string
	Name  string
}

// MissingNested returns every user in incoming whose Email is not already in
// existing.
//
// This is the version that gets written first, and it is not stupid. It reads
// exactly like the sentence describing the job: for each incoming user, look
// through the existing ones, and keep it if nothing matched. No extra memory,
// no setup, nothing to get wrong.
func MissingNested(incoming, existing []User) []User {
	var out []User
	for _, want := range incoming {
		found := false
		for _, have := range existing {
			if have.Email == want.Email {
				found = true
				break
			}
		}
		if !found {
			out = append(out, want)
		}
	}
	return out
}

// MissingSet answers the same question by building a set of the emails first.
//
// The loop over existing runs once, up front, instead of once per incoming
// user. After that each lookup is a single hash rather than a scan.
func MissingSet(incoming, existing []User) []User {
	seen := make(map[string]struct{}, len(existing))
	for _, have := range existing {
		// struct{}{} rather than true: the value is never read, and an empty
		// struct occupies no space, so the map stores only its keys.
		seen[have.Email] = struct{}{}
	}
	var out []User
	for _, want := range incoming {
		if _, ok := seen[want.Email]; !ok {
			out = append(out, want)
		}
	}
	return out
}
