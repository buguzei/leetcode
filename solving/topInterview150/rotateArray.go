package topInterview150

func RotateArray(nums []int, k int) {
	k = k % len(nums)

	numsMap := make(map[int]int, len(nums))

	var target int
	for i := range nums {
		if i+k <= len(nums)-1 {
			target = i + k
		} else {
			target = (i + k) - (len(nums))
		}

		if _, ok := numsMap[i]; ok {
			numsMap[target] = nums[target]

			nums[target] = numsMap[i]
		} else {
			numsMap[target] = nums[target]

			nums[target] = nums[i]
		}
	}
}
