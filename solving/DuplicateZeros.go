package solving

func DuplicateZeros(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			subSlice := nums[i+1:]
			for j := len(subSlice) - 1; j != 0; j-- {
				subSlice[j-1], subSlice[j] = subSlice[j], subSlice[j-1]
			}
			nums[i+1] = 0
			i++

		}
	}

	return nums
}
