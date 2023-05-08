package find_pivot_index

func pivotIndex(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == sum-leftSum-num {
			return i
		}

		leftSum += num
	}

	return -1
}
