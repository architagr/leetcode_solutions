// Package renumber moves days in the challenge queue and repairs every
// published-facing artifact that carries a day number.
//
// The day number is not private to queue.yaml. It is printed on each
// hero image, written into the header of five post files, and cited in
// prose and cross-reference links throughout a folder's write-ups — 435
// references across the queue as this package was written, in seven
// different file types. Moving a day is therefore a repo-wide edit, and
// doing it by hand is how references go stale.
package renumber

import (
	"fmt"
	"regexp"
	"strconv"
)

// dayRefRe matches every form a day reference takes: the post header
// "Day 12/365", the hero alt text "![Day 12](HERO.png)", a
// cross-reference link "[Day 14: Title](url)", and bare prose such as
// "the same lazy growth as Day 14's".
//
// All four are the same two tokens, so one pattern covers them and
// there is no form-specific branch to forget to update.
var dayRefRe = regexp.MustCompile(`Day (\d+)`)

// RewriteDayRefs replaces every day reference in content according to
// dayByOldDay, leaving references to unmapped days untouched.
//
// Replacement is simultaneous, which is the whole correctness argument:
// ReplaceAllStringFunc scans the input once and appends each replacement
// to a separate buffer, so output is never rescanned. A naive loop of
// per-day string replacements would chain — mapping 20→30 and 30→45
// would carry the original day 20 all the way to 45.
func RewriteDayRefs(content string, dayByOldDay map[int]int) string {
	return dayRefRe.ReplaceAllStringFunc(content, func(match string) string {
		old, err := strconv.Atoi(match[len("Day "):])
		if err != nil {
			return match
		}
		next, ok := dayByOldDay[old]
		if !ok {
			return match
		}
		return fmt.Sprintf("Day %d", next)
	})
}
