package main

import (
	"container/heap"
	"fmt"
)

type node struct {
	x int64
	y int64
}

func (n node) value() int64 { return n.x * n.y }

type nodeHeap []node

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].value() > h[j].value() }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{}) {
	*h = append(*h, x.(node))
}
func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func reverseDigits(n int64) int64 {
	reversed := int64(0)
	for n > 0 {
		digit := n % 10
		n /= 10
		reversed = reversed*10 + digit
	}
	return reversed
}

func isPalindrome(n int64) bool {
	return n == reverseDigits(n)
}

func main() {
	h := nodeHeap{node{999, 999}}
	set := make(map[node]struct{})
	for h.Len() > 0 {
		elem := heap.Pop(&h).(node)
		if isPalindrome(elem.value()) {
			fmt.Println(elem.value())
			return
		}
		if elem.x > 100 && elem.x-1 >= elem.y {
			nextNode := node{elem.x - 1, elem.y}
			if _, exists := set[nextNode]; !exists {
				heap.Push(&h, nextNode)
				set[nextNode] = struct{}{}
			}
		}
		if elem.y > 100 {
			nextNode := node{elem.x, elem.y - 1}
			if _, exists := set[nextNode]; !exists {
				heap.Push(&h, node{elem.x, elem.y - 1})
				set[nextNode] = struct{}{}
			}
		}
	}
}
