package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return isArrayPalindrome(arr)
}

func isArrayPalindrome(arr []int) bool {
	i, j := 0, len(arr)-1

	for i <= j {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
