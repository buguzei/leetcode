package topInterview150

func MajorityElement(nums []int) int {
	numsMap := make(map[int]int, len(nums))

	for _, el := range nums {
		if _, ok := numsMap[el]; ok {
			numsMap[el]++
		} else {
			numsMap[el] = 1
		}
	}

	var maxElem int
	for key, val := range numsMap {
		if numsMap[maxElem] < val {
			maxElem = key
		}
	}

	return maxElem
}
