package main

import "fmt"

func main() {
	for a := int64(1); a < 1000; a++ {
		for b := int64(1); a+b < 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
