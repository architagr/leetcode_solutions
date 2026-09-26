package whoisthecommonmanager

import "testing"

var company, staff = org(50000, 3, 8, 7) // 50,000 employees
var numbered = Number(company)

var shapes = []struct {
	name string
	ids  []int
}{
	{"pair", pick(staff, 2, 11)},             // two people, anywhere
	{"team_20", oneTeam(staff, 20, 12)},      // a team meeting
	{"allhands_200", pick(staff, 200, 13)},   // a cross-org review
	{"allhands_2000", pick(staff, 2000, 14)}, // an all-hands invite
}

func BenchmarkManagerByPaths(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ManagerByPaths(company, s.ids)
			}
		})
	}
}

func BenchmarkManagerByCount(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ManagerByCount(company, s.ids)
			}
		})
	}
}

func BenchmarkManagerByNumbers(b *testing.B) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ManagerByNumbers(company, numbered, s.ids)
			}
		})
	}
}

// Numbering is paid once per change to the chart, not per question.
func BenchmarkNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Number(company)
	}
}
