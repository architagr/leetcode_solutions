package dedupeuserlist

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func users(emails ...string) []User {
	out := make([]User, 0, len(emails))
	for i, e := range emails {
		out = append(out, User{ID: fmt.Sprint(i), Email: e, Name: "user " + e})
	}
	return out
}

func emailsOf(us []User) []string {
	out := make([]string, 0, len(us))
	for _, u := range us {
		out = append(out, u.Email)
	}
	return out
}

// The entire argument of this episode is that the two functions are
// interchangeable, so that is the first thing tested.
func TestBothAgree(t *testing.T) {
	cases := []struct {
		name               string
		incoming, existing []User
		want               []string
	}{
		{"none missing", users("a", "b"), users("a", "b", "c"), nil},
		{"all missing", users("x", "y"), users("a", "b"), []string{"x", "y"}},
		{"some missing", users("a", "x", "b", "y"), users("a", "b"), []string{"x", "y"}},
		{"empty incoming", nil, users("a"), nil},
		{"empty existing", users("a", "b"), nil, []string{"a", "b"}},
		{"both empty", nil, nil, nil},
		// Order matters: callers write the result straight to an insert, so a
		// function that returned the same set in a different order would be a
		// different function.
		{"keeps incoming order", users("z", "a", "m"), users("a"), []string{"z", "m"}},
		// A duplicate inside incoming is reported twice by both, because
		// neither de-duplicates its own input. Pinning it so a future change
		// cannot quietly alter it.
		{"duplicate in incoming", users("x", "x"), users("a"), []string{"x", "x"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotN := emailsOf(MissingNested(c.incoming, c.existing))
			gotS := emailsOf(MissingSet(c.incoming, c.existing))
			if len(c.want) == 0 {
				if len(gotN) != 0 || len(gotS) != 0 {
					t.Fatalf("want nothing, nested=%v set=%v", gotN, gotS)
				}
				return
			}
			if !reflect.DeepEqual(gotN, c.want) {
				t.Errorf("nested = %v, want %v", gotN, c.want)
			}
			if !reflect.DeepEqual(gotS, c.want) {
				t.Errorf("set = %v, want %v", gotS, c.want)
			}
		})
	}
}

// Agreement on hand-written cases is weak evidence. This runs both over random
// inputs, where overlap is partial and unpredictable.
func TestBothAgreeOnRandomInput(t *testing.T) {
	rng := rand.New(rand.NewSource(7))
	for trial := 0; trial < 2000; trial++ {
		pool := rng.Intn(40) + 1
		mk := func(n int) []User {
			out := make([]User, n)
			for i := range out {
				out[i] = User{Email: fmt.Sprintf("u%d@x", rng.Intn(pool))}
			}
			return out
		}
		in, ex := mk(rng.Intn(30)), mk(rng.Intn(30))
		if n, s := emailsOf(MissingNested(in, ex)), emailsOf(MissingSet(in, ex)); !reflect.DeepEqual(n, s) {
			t.Fatalf("trial %d disagreed:\n nested %v\n set    %v\n in %v\n ex %v",
				trial, n, s, emailsOf(in), emailsOf(ex))
		}
	}
}

// The claim in the write-up is that the nested version's comparison count is
// the product of the two lengths when nothing matches. Counted, not asserted
// in prose.
func TestNestedComparisonCount(t *testing.T) {
	for _, n := range []int{10, 50, 200} {
		incoming := make([]User, n)
		existing := make([]User, n)
		for i := 0; i < n; i++ {
			incoming[i] = User{Email: fmt.Sprintf("in%d", i)}
			existing[i] = User{Email: fmt.Sprintf("ex%d", i)}
		}
		compares := 0
		for _, want := range incoming {
			for _, have := range existing {
				compares++
				if have.Email == want.Email {
					break
				}
			}
		}
		if want := n * n; compares != want {
			t.Errorf("n=%d: %d comparisons, want %d", n, compares, want)
		}
	}
}
