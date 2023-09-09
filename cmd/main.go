package main

import (
	"fmt"
	"github.com/buguzei/leetcode/solving"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	fmt.Println(solving.MergeSortedArray(nums1, nums2, 3, 3))
}
