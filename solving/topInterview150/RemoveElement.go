package topInterview150

import "fmt"

func RemoveElement(nums []int, val int) int {
	var counter int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			counter++
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	fmt.Println(nums)

	return counter
}
