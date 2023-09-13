package arrays101

func MaxConsecutiveOnes(nums []int) int {
	var currentCounter int
	var maxCounter int

	for i, el := range nums {
		if el == 1 {
			currentCounter++

			if i != len(nums)-1 {
				continue
			}
		}

		if currentCounter > maxCounter {
			maxCounter = currentCounter
		}

		currentCounter = 0
	}

	return maxCounter
}
