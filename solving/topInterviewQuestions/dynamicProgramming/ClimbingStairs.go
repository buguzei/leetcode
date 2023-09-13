package dynamicProgramming

func ClimbingStairs(n int) int {
	nums := make([]int, n+1)

	nums[0] = 1
	nums[1] = 1

	for i := 2; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums[len(nums)-1]
}
