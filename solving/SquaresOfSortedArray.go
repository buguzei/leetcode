package solving

import "fmt"

func SquaresOfSortedArray(nums []int) []int {
	for i, _ := range nums {
		nums[i] *= nums[i]
	}
	fmt.Println(nums)

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
