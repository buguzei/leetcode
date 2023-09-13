package main

import (
	"fmt"
	"github.com/buguzei/leetcode/solving/topInterviewQuestions/dynamicProgramming"
)

func main() {
	nums := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(dynamicProgramming.BestTimetoBuyandSellStock(nums))
}
