func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxConsecutive := 0
	for _, n := range nums {
		if (n == 1) {
			count += 1
		} else {
			count = 0
		}

		if (count > maxConsecutive) {
			maxConsecutive = count
		}
	}
	return maxConsecutive
}
