package topInterview150

func JumpGame(nums []int) bool {
	maxi := 0
	for i := 0; i < len(nums); i++ {
		if i > maxi {
			return false
		}
		if i+nums[i] > maxi {
			maxi = i + nums[i]
		}
	}
	return true
}
