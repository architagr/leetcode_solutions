package sentencesimilarityiii

import (
	"fmt"
	"testing"
)

type TestCase struct {
	s1, s2   string
	expected bool
}

func TestAreSentencesSimilar(t *testing.T) {
	testCases := []TestCase{
		{
			s1:       "Hello Jane",
			s2:       "Hello my name is Jane",
			expected: true,
		},
		{
			s1:       "Frogs are cool",
			s2:       "Frog cool",
			expected: false,
		},
		{
			s1:       "My name is Haley",
			s2:       "My Haley",
			expected: true,
		},
		{
			s1:       "of",
			s2:       "A lot of words",
			expected: false,
		},
		{
			s1:       "Eating right now",
			s2:       "Eating",
			expected: true,
		},
		{
			s1:       "Eating right now",
			s2:       "now",
			expected: true,
		},
		{
			s1:       "A A AAa",
			s2:       "A AAa",
			expected: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t1 *testing.T) {
			if got := areSentencesSimilar(tc.s1, tc.s2); got != tc.expected {
				t1.Errorf("expected %t but got %t", tc.expected, got)
			}
		})
	}
}
