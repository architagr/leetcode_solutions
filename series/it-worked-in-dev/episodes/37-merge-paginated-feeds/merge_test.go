package mergepaginatedfeeds

import (
	"math/rand"
	"slices"
	"testing"
)

// following is k feeds of up to per posts each, newest first. Accounts post
// at very different rates: a few post every few minutes, most a few times a
// week, so the newest posts come from a handful of feeds.
func following(k, per int, seed int64) [][]Post {
	r := rand.New(rand.NewSource(seed))
	now := int64(1_790_000_000_000)
	var id int64
	feeds := make([][]Post, k)
	for f := range feeds {
		gap := int64(60_000 * (1 + r.ExpFloat64()*600)) // mean gap between this account's posts
		at := now - r.Int63n(gap)
		posts := make([]Post, per)
		for i := range posts {
			id++
			posts[i] = Post{id, at}
			at -= 1 + int64(r.ExpFloat64()*float64(gap))
		}
		feeds[f] = posts
	}
	return feeds
}

var funcs = map[string]func([][]Post, int) []Post{
	"sort": PageBySort, "fold": PageByFolding, "heap": PageByHeap,
}

func TestAllAgree(t *testing.T) {
	cases := [][][]Post{nil, {}, {{}}, {{}, {{1, 5}}, {}}, following(1, 3, 1), following(7, 20, 2), following(50, 100, 3)}
	for _, s := range shapes {
		cases = append(cases, s.feeds)
	}
	for _, c := range cases {
		for _, n := range []int{1, 50, 1000} {
			want := PageBySort(c, n)
			for name, fn := range funcs {
				if got := fn(c, n); !slices.Equal(got, want) {
					t.Fatalf("%d feeds, n=%d: %s disagrees with sort", len(c), n, name)
				}
			}
		}
	}
}

func TestFeedsUntouched(t *testing.T) {
	feeds := following(20, 50, 9)
	before := make([][]Post, len(feeds))
	for i := range feeds {
		before[i] = slices.Clone(feeds[i])
	}
	for _, fn := range funcs {
		fn(feeds, 50)
	}
	for i := range feeds {
		if !slices.Equal(feeds[i], before[i]) {
			t.Fatal("a feed was modified")
		}
	}
}

// What each version touches.
func TestWork(t *testing.T) {
	for _, s := range shapes {
		total := 0
		for _, f := range s.feeds {
			total += len(f)
		}
		copied := 0 // posts the fold copies, summed over every merge
		size := 0
		for _, f := range s.feeds {
			size += len(f)
			copied += size
		}
		t.Logf("%-8s %5d feeds, %7d posts, page %4d  |  fold copies %11d posts",
			s.name, len(s.feeds), total, s.n, copied)
	}
}
