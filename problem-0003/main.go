package main

import (
	"fmt"
	"math"
)

func largestPrimeFactor(n int64) int64 {
	var maxFactor int64 = -1
	for factor := int64(2); factor < int64(math.Sqrt(float64(n))); factor++ {
		if n%factor == 0 {
			n /= factor
			maxFactor = n
		}
	}
	return maxFactor
}

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}
