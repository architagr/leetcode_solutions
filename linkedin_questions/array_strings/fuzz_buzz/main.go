package fuzzbuzz

import "fmt"

func fizzBuzz(n int) []string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		by3 := (i+1)%3 == 0
		by5 := (i+1)%5 == 0
		if by5 && by3 {
			arr[i] = "FizzBuzz"
		} else if by3 {
			arr[i] = "Fizz"
		} else if by5 {
			arr[i] = "Buzz"
		} else {
			arr[i] = fmt.Sprint(i + 1)
		}
	}
	return arr
}
