package queue

import (
	"testing"
	"time"
)

func ts(s string) *string { return &s }

func TestSelectNextPicksOldestUnpostedForThatDestination(t *testing.T) {
	q := &Queue{Entries: []Entry{
		// Deliberately out of day order: selection must go by day, not file order.
		{Day: 3, Number: 257},
		{Day: 1, Number: 104, PostedAt: map[string]*string{DestinationDiscord: ts("2026-09-01T09:00:00Z")}},
		{Day: 2, Number: 108, PostedAt: map[string]*string{DestinationLinkedInMain: ts("2026-09-02T09:00:00Z")}},
	}}

	got, ok := SelectNext(q, DestinationDiscord)
	if !ok {
		t.Fatal("SelectNext: ok = false, want true")
	}
	// Day 1 is already on discord; day 2 is only on linkedin, so it's next for discord.
	if got.Day != 2 {
		t.Errorf("SelectNext picked day %d, want 2", got.Day)
	}
}

func TestSelectNextEmptyAndFullyPosted(t *testing.T) {
	if _, ok := SelectNext(&Queue{}, DestinationDiscord); ok {
		t.Error("SelectNext on an empty queue = true, want false")
	}

	allPosted := &Queue{Entries: []Entry{
		{Day: 1, Number: 104, PostedAt: map[string]*string{DestinationDiscord: ts("2026-09-01T09:00:00Z")}},
	}}
	if _, ok := SelectNext(allPosted, DestinationDiscord); ok {
		t.Error("SelectNext with everything posted = true, want false")
	}
}

func TestSelectNextTreatsExplicitNullAsUnposted(t *testing.T) {
	q := &Queue{Entries: []Entry{
		{Day: 1, Number: 104, PostedAt: map[string]*string{DestinationDiscord: nil}},
	}}
	got, ok := SelectNext(q, DestinationDiscord)
	if !ok || got.Day != 1 {
		t.Errorf("SelectNext = (%+v, %v), want day 1 selected", got, ok)
	}
}

func TestPostedOnFindsAnEntrySentThatUTCDay(t *testing.T) {
	q := &Queue{Entries: []Entry{
		{Day: 1, Number: 104, PostedAt: map[string]*string{DestinationDiscord: ts("2026-09-03T06:16:03Z")}},
		{Day: 2, Number: 108, PostedAt: map[string]*string{DestinationDiscord: ts("2026-09-04T05:31:13Z")}},
		{Day: 3, Number: 257},
	}}

	// A second run later the same UTC day must see day 2 as today's post.
	got, ok := PostedOn(q, DestinationDiscord, time.Date(2026, 9, 4, 23, 59, 0, 0, time.UTC))
	if !ok {
		t.Fatal("PostedOn: ok = false, want true")
	}
	if got.Day != 2 {
		t.Errorf("PostedOn found day %d, want 2", got.Day)
	}

	// The next UTC day is clear again, even a minute past midnight.
	if _, ok := PostedOn(q, DestinationDiscord, time.Date(2026, 9, 5, 0, 1, 0, 0, time.UTC)); ok {
		t.Error("PostedOn on a fresh day = true, want false")
	}
}

func TestPostedOnComparesInUTCAndIgnoresOtherDestinations(t *testing.T) {
	q := &Queue{Entries: []Entry{
		{Day: 1, Number: 104, PostedAt: map[string]*string{DestinationLinkedInMain: ts("2026-09-04T05:00:00Z")}},
		{Day: 2, Number: 108, PostedAt: map[string]*string{DestinationDiscord: ts("bogus")}},
	}}

	// Another destination's post that day must not block discord.
	if _, ok := PostedOn(q, DestinationDiscord, time.Date(2026, 9, 4, 6, 0, 0, 0, time.UTC)); ok {
		t.Error("PostedOn matched a different destination, want false")
	}

	// A now in a non-UTC zone still compares by UTC calendar day: 2026-09-05
	// 04:00 +05:30 is 2026-09-04 22:30 UTC.
	q.Entries[1].PostedAt[DestinationDiscord] = ts("2026-09-04T05:31:13Z")
	ist := time.FixedZone("IST", 5*60*60+30*60)
	if _, ok := PostedOn(q, DestinationDiscord, time.Date(2026, 9, 5, 4, 0, 0, 0, ist)); !ok {
		t.Error("PostedOn with a non-UTC now = false, want true")
	}
}
