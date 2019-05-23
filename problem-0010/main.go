package main

import (
	"fmt"
	"github.com/priyendra/golang-euler/common"
)

func main() {
	sum := int64(0)
	for i := int64(0); i < 2000000; i++ {
		if common.IsPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
