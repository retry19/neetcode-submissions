func replaceElements(arr []int) []int {
	maxSoFar := -1
	for i := len(arr)-1; i >= 0; i-- {
		arr[i], maxSoFar = maxSoFar, max(arr[i], maxSoFar)
	}
	return arr
}
