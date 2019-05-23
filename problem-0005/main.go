package main

import (
	"fmt"
	"math"
)

func pow(a, b int) int {
	answer := 1
	for i := 0; i < b; i++ {
		answer *= a
	}
	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPrime(n int) bool {
	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func allPrimesUpTo(n int) []int {
	answer := []int{2}
	for i := 3; i <= n; i++ {
		if isPrime(i) {
			answer = append(answer, i)
		}
	}
	return answer
}

func main() {
	N := 20
	primes := allPrimesUpTo(N)
	powers := make([]int, len(primes))
	for i := 4; i <= N; i++ {
		for j := 0; j < len(primes); j++ {
			if i%primes[j] == 0 {
				power := 0
				for i2 := i; i2 != 0; i2 /= primes[j] {
					power++
				}
				powers[j] = max(powers[j], power-1)
			}
		}
	}
	answer := 1
	for i := 0; i < len(primes); i++ {
		answer *= pow(primes[i], powers[i])
	}
	fmt.Println(answer)
}
