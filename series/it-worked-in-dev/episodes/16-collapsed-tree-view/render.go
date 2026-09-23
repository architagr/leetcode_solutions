package collapsedtreeview

import (
	"strconv"
	"strings"
)

// The collapsed view does not show an ID. It shows a preview line, and
// building one costs something: trimming the body, escaping it, formatting a
// relative timestamp, counting the replies underneath.
//
// Preview stands in for that. It is deliberately not free, and it is the same
// cost in both versions below, so the only thing the benchmark compares is how
// many times each one calls it.
func Preview(c *Comment) string {
	var b strings.Builder
	b.Grow(48)
	b.WriteString(c.ID)
	b.WriteString(" · ")
	b.WriteString(strconv.Itoa(len(c.Replies)))
	b.WriteString(" replies · ")
	for i := 0; i < 3; i++ {
		b.WriteString(c.ID)
	}
	return b.String()
}

// PreviewsByDepthMap keeps one entry per depth and lets the last writer win,
// which means every comment in the thread gets a preview built for it and all
// but one per level is thrown away.
func PreviewsByDepthMap(root *Comment) ([]string, int) {
	var out []string
	built := 0
	var walk func(c *Comment, depth int)
	walk = func(c *Comment, depth int) {
		if c == nil {
			return
		}
		if len(out) == depth {
			out = append(out, "")
		}
		built++
		out[depth] = Preview(c)
		for _, r := range c.Replies {
			walk(r, depth+1)
		}
	}
	walk(root, 0)
	return out, built
}

// PreviewsByFirstArrival walks replies newest first, so the comment that wins
// a level is the first one reached there - and no preview is ever built for a
// comment that is about to be overwritten.
func PreviewsByFirstArrival(root *Comment) ([]string, int) {
	var out []string
	built := 0
	var walk func(c *Comment, depth int)
	walk = func(c *Comment, depth int) {
		if c == nil {
			return
		}
		if len(out) == depth {
			built++
			out = append(out, Preview(c))
		}
		for i := len(c.Replies) - 1; i >= 0; i-- {
			walk(c.Replies[i], depth+1)
		}
	}
	walk(root, 0)
	return out, built
}
