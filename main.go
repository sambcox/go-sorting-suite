package main

import (
	"fmt"
	"github.com/sambcox/go-sorting-suite/bubbleSort"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	b := bubbleSort.BubbleSort{}
	fmt.Println(b.Sort(arr))
}
