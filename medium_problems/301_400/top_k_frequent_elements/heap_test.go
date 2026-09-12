package top_k_frequent_elements

import (
	"math/rand"
	"sort"
	"testing"
)

// The heap holds the top k by frequency, in no particular order, so the
// check is on the set of values rather than the sequence.
func sorted(v []int) []int {
	out := append([]int(nil), v...)
	sort.Ints(out)
	return out
}

func eq(a, b []int) bool {
	a, b = sorted(a), sorted(b)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTopKFrequentEdgeCases(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"k equals distinct count", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"k is 1", []int{4, 4, 4, 9, 9, 2}, 1, []int{4}},
		{"all identical", []int{7, 7, 7}, 1, []int{7}},
		{"negatives", []int{-1, -1, -2, -2, -2, 5}, 2, []int{-1, -2}},
		{"clear frequency order", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := TopKFrequent(c.nums, c.k); !eq(got, c.want) {
				t.Errorf("TopKFrequent(%v, %d) = %v, want %v", c.nums, c.k, got, c.want)
			}
		})
	}
}

// Counting and sorting by frequency is the definition of the answer, so it
// makes a reference. Inputs are generated so the kth and (k+1)th frequencies
// differ, because otherwise several answers are equally correct and the
// problem guarantees that does not happen.
func TestTopKFrequentAgainstSort(t *testing.T) {
	r := rand.New(rand.NewSource(3))
	for trial := 0; trial < 2000; trial++ {
		counts := map[int]int{}
		distinct := 1 + r.Intn(6)
		nums := []int{}
		used := map[int]bool{}
		for v := 0; len(counts) < distinct; v++ {
			if used[v] {
				continue
			}
			used[v] = true
			c := 1 + r.Intn(8)
			counts[v] = c
			for i := 0; i < c; i++ {
				nums = append(nums, v)
			}
		}
		type fc struct{ v, c int }
		all := []fc{}
		for v, c := range counts {
			all = append(all, fc{v, c})
		}
		sort.Slice(all, func(i, j int) bool { return all[i].c > all[j].c })

		k := 1 + r.Intn(len(all))
		if k < len(all) && all[k-1].c == all[k].c {
			continue // ambiguous boundary; the problem excludes these
		}
		want := make([]int, k)
		for i := 0; i < k; i++ {
			want[i] = all[i].v
		}
		r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
		if got := TopKFrequent(nums, k); !eq(got, want) {
			t.Fatalf("counts %v k=%d: got %v, want %v", counts, k, got, want)
		}
	}
}
