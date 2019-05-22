package main

import "fmt"

func main() {
	x := 1
	y := 2
	sum := 0
	for y < 4000000 {
		if y%2 == 0 {
			sum = sum + y
		}
		tmp := x
		x = y
		y = tmp + x
	}
	fmt.Println(sum)
}
