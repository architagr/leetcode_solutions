package interleaving_string

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputS1, inputS2, inputS3 string
	expected                  bool
}

var testCases []TestCase = []TestCase{
	/*1*/ {inputS1: "", inputS2: "dbbca", inputS3: "dbbca", expected: true},
	/*2*/ {inputS1: "", inputS2: "dbbca", inputS3: "aadbbbaccc", expected: false},
	/*3*/ {inputS1: "aabcc", inputS2: "", inputS3: "aabcc", expected: true},
	/*4*/ {inputS1: "aabcc", inputS2: "", inputS3: "aadbbbaccc", expected: false},
	/*5*/ {inputS1: "", inputS2: "", inputS3: "", expected: true},
	/*6*/ {inputS1: "a", inputS2: "b", inputS3: "a", expected: false},
	/*7*/ {inputS1: "aabcc", inputS2: "dbbca", inputS3: "aadbbcbcac", expected: true},
	/*8*/ {inputS1: "aabcc", inputS2: "dbbca", inputS3: "aadbbbaccc", expected: false},
	/*9*/ {inputS1: "aa", inputS2: "ba", inputS3: "baaa", expected: true},
	/*10*/ {inputS1: "aa", inputS2: "ab", inputS3: "abaa", expected: true},
	/*11*/ {inputS1: "db", inputS2: "b", inputS3: "cbb", expected: false},
	/*12*/ {inputS1: "ab", inputS2: "bc", inputS3: "babc", expected: true},
	/*13*/ {inputS1: "aabc", inputS2: "abad", inputS3: "aabadabc", expected: true},
}

func TestCheckInterleavingString(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := isInterleave(testCase.inputS1, testCase.inputS2, testCase.inputS3)
			if got != testCase.expected {
				test.Fatalf("testing %d expected %t but got %t", (i + 1), testCase.expected, got)
			}
		})
	}
}
