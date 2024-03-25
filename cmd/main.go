package main

import (
	"fmt"
	"github.com/buguzei/leetcode/solving/topInterview150"
)

func main() {
	gas := []int{5, 8, 2, 8}
	cost := []int{6, 5, 6, 6}

	fmt.Println(topInterview150.GasStation(gas, cost))
}
