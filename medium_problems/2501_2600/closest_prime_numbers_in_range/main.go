package closestprimenumbersinrange

import "math"

func closestPrimes(left int, right int) []int {
	primeNumbers := []int{}

	// Find all prime numbers in the given range
	for candidate := left; candidate <= right; candidate++ {
		if candidate%2 == 0 && candidate > 2 {
			continue
		}

		if isPrime(candidate) {
			// If a twin prime (or [2, 3]) is found, return immediately
			if len(primeNumbers) > 0 && candidate <= primeNumbers[len(primeNumbers)-1]+2 {
				return []int{
					primeNumbers[len(primeNumbers)-1],
					candidate,
				}
			}
			primeNumbers = append(primeNumbers, candidate)
		}
	}

	// If fewer than 2 primes exist, return {-1, -1}
	if len(primeNumbers) < 2 {
		return []int{-1, -1}
	}

	// Find the closest prime pair
	closestPair := []int{-1, -1}
	minDifference := math.MaxInt

	for index := 1; index < len(primeNumbers); index++ {
		difference := primeNumbers[index] - primeNumbers[index-1]

		if difference < minDifference {
			minDifference = difference
			closestPair = []int{primeNumbers[index-1], primeNumbers[index]}
		}
	}

	return closestPair
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	maxNum := int(math.Sqrt(float64(num)))
	for divisor := 2; divisor <= maxNum; divisor++ {
		if num%divisor == 0 {
			return false
		}
	}
	return true
}
