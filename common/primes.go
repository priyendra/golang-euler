package common

import "math"

func IsPrime(n int) bool {
	max := int(math.Sqrt(float64(n)))
	if n < 2 {
		return false
	}
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
