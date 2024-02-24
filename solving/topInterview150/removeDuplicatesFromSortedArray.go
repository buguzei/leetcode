package topInterview150

func RemoveDuplicatesFromSortedArray(nums []int) int {
	before := -101
	for i := 0; i < len(nums); i++ {
		if nums[i] == before {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
		before = nums[i]
	}

	return len(nums)
}
