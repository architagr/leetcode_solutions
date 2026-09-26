package rangequerywithoutscan

import "testing"

var cat = catalog(200000, 7)       // 200,000 products, prices 0 to 19,999.90
var imported = sortedImport(20000) // 20,000 products bulk-imported in price order

var shapes = []struct {
	name   string
	root   *Node
	lo, hi int
}{
	{"narrow_20", cat, 1000000, 1000190},         // ₹10,000.00 to ₹10,001.90: 20 products
	{"band_2k", cat, 1000000, 1019990},           // a ₹200 band: 2,000 products
	{"wide_100k", cat, 500000, 1499990},          // half the catalogue
	{"everything", cat, 0, 1999990},              // the whole catalogue
	{"import_low_20", imported, 0, 190},          // bulk import, the 20 cheapest
	{"import_high_20", imported, 199800, 199990}, // bulk import, the 20 dearest
}

func run(b *testing.B, fn func(*Node, int, int) []int) {
	for _, s := range shapes {
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(s.root, s.lo, s.hi)
			}
		})
	}
}

func BenchmarkRangeByScan(b *testing.B)    { run(b, RangeByScan) }
func BenchmarkRangeByPruning(b *testing.B) { run(b, RangeByPruning) }
