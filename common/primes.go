package common

import "math"

func IsPrime(n int64) bool {
	max := int64(math.Sqrt(float64(n)))
	if n < 2 {
		return false
	}
	for i := int64(2); i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
