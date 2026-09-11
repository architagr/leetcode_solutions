package palindrome_linked_list

import "testing"

// buildList turns the array form the examples use into a linked list.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want bool
	}{
		{name: "example 1", head: []int{1, 2, 2, 1}, want: true},
		{name: "example 2", head: []int{1, 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(buildList(tt.head)); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
