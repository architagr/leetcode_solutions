package accountsmerge

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	accounts [][]string
	expected [][]string
}

func TestAccountsMerge(t *testing.T) {
	testcases := []testcase{
		{
			accounts: [][]string{
				{"David", "David0@m.co", "David1@m.co"},
				{"David", "David3@m.co", "David4@m.co"},
				{"David", "David4@m.co", "David5@m.co"},
				{"David", "David2@m.co", "David3@m.co"},
				{"David", "David1@m.co", "David2@m.co"},
			},
			expected: [][]string{
				{"David", "David0@m.co", "David1@m.co", "David2@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
			},
		},
		{
			accounts: [][]string{
				{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
				{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
				{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
				{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
			},
			expected: [][]string{
				{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
				{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe2@m.co", "Gabe3@m.co", "Gabe4@m.co"},
				{"Kevin", "Kevin2@m.co", "Kevin4@m.co"},
			},
		},
		{
			accounts: [][]string{
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
			expected: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
		},
		{
			accounts: [][]string{
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
				{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
				{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
				{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
				{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
			},
			expected: [][]string{
				{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
				{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
				{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
				{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := accountsMerge(tc.accounts)
			sort.Slice(got, func(i, j int) bool {
				return got[i][0] < got[j][0]
			})
			sort.Slice(tc.expected, func(i, j int) bool {
				return tc.expected[i][0] < tc.expected[j][0]
			})
			assert.Equal(t, tc.expected, got)
		})
	}
}
