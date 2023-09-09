package solving

func MaxConsecutiveOnes(slc []int) int {
	var currentCounter int
	var maxCounter int

	for i, el := range slc {
		if el == 1 {
			currentCounter++

			if i != len(slc)-1 {
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
