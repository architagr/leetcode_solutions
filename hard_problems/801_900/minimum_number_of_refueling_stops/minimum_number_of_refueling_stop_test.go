package minimum_number_of_refueling_stops

import (
	"fmt"
	"testing"
)

type TestCase struct {
	target, startFuel int
	stations          [][]int
	expected          int
}

var testCases []TestCase = []TestCase{
	{
		target:    1,
		startFuel: 1,
		stations:  [][]int{},
		expected:  0,
	},
	{
		target:    100,
		startFuel: 1,
		stations:  [][]int{{10, 100}},
		expected:  -1,
	},
	{
		target:    100,
		startFuel: 10,
		stations:  [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}},
		expected:  2,
	},
	{
		target:    100,
		startFuel: 1,
		stations:  [][]int{},
		expected:  -1,
	},
	{
		target:    100,
		startFuel: 50,
		stations:  [][]int{{25, 30}},
		expected:  -1,
	},
	{
		target:    1000,
		startFuel: 83,
		stations:  [][]int{{25, 27}, {36, 187}, {140, 186}, {378, 6}, {492, 202}, {517, 89}, {579, 234}, {673, 86}, {808, 53}, {954, 49}},
		expected:  -1,
	},
	{
		target:    1000,
		startFuel: 299,
		stations:  [][]int{{13, 21}, {26, 115}, {100, 47}, {225, 99}, {299, 141}, {444, 198}, {608, 190}, {636, 157}, {647, 255}, {841, 123}},
		expected:  4,
	},
}

func TestMinRefuelStops(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := MinRefuelStops(testCase.target, testCase.startFuel, testCase.stations)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
