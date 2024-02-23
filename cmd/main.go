package main

import (
	"fmt"
	"github.com/buguzei/leetcode/solving/graph"
)

func main() {
	nums := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(graph.NumberOfProvinces(nums))
}
