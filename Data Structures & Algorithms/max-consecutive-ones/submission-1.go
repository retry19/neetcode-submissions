func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxConsecutive := 0
	for _, n := range nums {
		if (n == 1) {
			count += 1
		} else {
			count = 0
		}

		maxConsecutive = max(count, maxConsecutive)
	}
	return maxConsecutive
}
