package removeduplicatesfromsortedarray

import (
	"fmt"
	"testing"
)

type testCaseData struct{
    nums []int
    expectedLength int
}

var testCases []testCaseData = []testCaseData{
    {nums:[]int{1,1,2}, expectedLength: 2 },
    {nums:[]int{0,0,1,1,1,2,2,3,3,4}, expectedLength: 5 },
    {nums:[]int{}, expectedLength: 0 },
}
func TestRemoveDuplicate(t *testing.T){

    for i, testCase := range testCases{
        t.Run(fmt.Sprintf("Test case no %d", (i+1)), func(t *testing.T) {
            got := removeDuplicates(testCase.nums)   
            if got != testCase.expectedLength{
                t.Errorf("error expected length %d and got %d\n", testCase.expectedLength, got)
            }
        })
    }
}