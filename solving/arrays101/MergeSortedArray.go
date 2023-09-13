package arrays101

func MergeSortedArray(nums1, nums2 []int, m, n int) []int {
	mergeNums1 := nums1[:m]
	mergeNums2 := nums2[:n]

	mergeNums := append(mergeNums1, mergeNums2...)

	for i := 0; i < len(mergeNums)-1; i++ {
		for j := 0; j < len(mergeNums)-i-1; j++ {
			if mergeNums[j+1] < mergeNums[j] {
				mergeNums[j], mergeNums[j+1] = mergeNums[j+1], mergeNums[j]
			}
		}
	}

	return mergeNums
}
