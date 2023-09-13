package arrays101

func FindNumbersWithEvenNumberOfDigits(nums []int) int {
	var counter int
	for _, el := range nums {
		var digit int
		for el != 0 {
			digit++
			el /= 10
		}

		if digit%2 == 0 {
			counter++
		}
	}

	return counter
}
