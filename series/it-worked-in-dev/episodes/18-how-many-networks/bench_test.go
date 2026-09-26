package howmanynetworks

import "testing"

// The shapes a mesh takes. teams_5k is the healthy one: many small networks,
// where asking every service for its reach costs almost nothing extra.
var shapes = []struct {
	name string
	mesh Mesh
}{
	{"isolated_5k", isolated(5000)},  // nothing linked: 5,000 networks of one
	{"teams_5k", teams(500, 10)},     // 500 teams, ten services each
	{"joined_5k", joined(500, 10)},   // the same 500 teams, plus 499 cross-team links
	{"mesh_500", mesh(500, 500, 7)},  // one small shop, everything linked
	{"mesh_5k", mesh(5000, 5000, 8)}, // one company-wide mesh
}

func run(b *testing.B, fn func(Mesh) int) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.mesh)
			}
		})
	}
}

func BenchmarkNetworksByReach(b *testing.B) { run(b, NetworksByReach) }
func BenchmarkNetworksBySweep(b *testing.B) { run(b, NetworksBySweep) }
