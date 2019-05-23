package main

import (
	"fmt"
	"github.com/priyendra/golang-euler/common"
)

func main() {
	n := 1
	index := 0
	for {
		if common.IsPrime(n) {
			index++
		}
		if index == 10001 {
			break
		}
		n++
	}
	fmt.Println(n)
}
