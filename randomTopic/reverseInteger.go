package randomTopic

import (
	"math"
)

func ReverseInteger(x int) int {
	var isPositive = true

	if x < 0 {
		x = -x
		isPositive = false
	}

	var digitAmount int

	testX := x

	for testX > 0 {
		digitAmount++
		testX /= 10
	}

	for i, j := digitAmount-1, 0; j < digitAmount/2; i, j = i-1, j+1 {
		if (x < 0 && isPositive) || (x > int(math.Pow(float64(2), float64(31))) || x < int(-math.Pow(float64(2), float64(31)))) {
			return 0
		}

		lPow := int(math.Pow10(i))
		rPow := int(math.Pow10(j))

		leftDigit := (x / lPow) % 10
		rightDigit := (x / rPow) % 10

		x = ((x/lPow)-(x/lPow%10)+rightDigit)*lPow + (x % lPow)
		x = ((x/rPow)-(x/rPow%10)+leftDigit)*rPow + (x % rPow)
	}

	if !isPositive {
		return -x
	}

	return x
}
