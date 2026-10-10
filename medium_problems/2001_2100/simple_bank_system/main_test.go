package simplebanksystem

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank1(t *testing.T) {
	b := Constructor([]int64{767, 653, 252, 849, 480, 187, 761, 243, 408, 385, 334, 732, 289, 886, 149, 320, 827, 111, 315, 155, 695, 110, 473, 585, 83, 936, 188, 818, 33, 984, 66, 549, 954, 761, 662, 212, 208, 215, 251, 792, 956, 261, 863, 374, 411, 639, 599, 418, 909, 208, 984, 602, 741, 302, 911, 616, 537, 422, 61, 746, 206, 396, 446, 661, 48, 156, 725, 662, 422, 624, 704, 143, 94, 702, 126, 76, 539, 83, 270, 717, 736, 393, 607, 895, 661})
	const deposit = "deposit"
	const transfer = "transfer"
	const withdraw = "withdraw"
	testcases := []struct {
		action   string
		args     []int64
		expected bool
	}{
		{
			action:   deposit,
			args:     []int64{68, 668},
			expected: true,
		},
		{
			action:   deposit,
			args:     []int64{25, 978},
			expected: true,
		},
		{
			action:   transfer,
			args:     []int64{8, 31, 924},
			expected: false,
		},
		{
			action:   transfer,
			args:     []int64{2, 6, 857},
			expected: false,
		},
		{
			action:   transfer,
			args:     []int64{20, 43, 59},
			expected: true,
		},
		{
			action:   deposit,
			args:     []int64{71, 307},
			expected: true,
		},
		{
			action:   transfer,
			args:     []int64{11, 46, 577},
			expected: false,
		},
		{
			action:   withdraw,
			args:     []int64{37, 377},
			expected: false,
		},
		{
			action:   deposit,
			args:     []int64{72, 835},
			expected: true,
		},
		{
			action:   withdraw,
			args:     []int64{82, 574},
			expected: false,
		},
		{
			action:   transfer,
			args:     []int64{67, 9, 939},
			expected: false,
		},
		{
			action:   transfer,
			args:     []int64{24, 49, 251},
			expected: true,
		},
	}
	for i, tc := range testcases {
		log.Print(i)
		t.Run(fmt.Sprint(i), func(tb *testing.T) {
			switch tc.action {
			case deposit:
				assert.Equal(tb, tc.expected, b.Deposit(int(tc.args[0]), tc.args[1]))
			case transfer:
				assert.Equal(tb, tc.expected, b.Transfer(int(tc.args[0]), int(tc.args[1]), tc.args[2]))
			default:
				assert.Equal(tb, tc.expected, b.Withdraw(int(tc.args[0]), tc.args[1]))
			}
		})
	}
}

// ["withdraw","transfer","transfer"]
// [[82,574],[67,9,939],[24,49,251]]

// false,false,true
