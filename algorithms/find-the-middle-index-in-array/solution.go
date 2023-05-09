package find_the_middle_index_in_array

func findMiddleIndex(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	leftSum := 0

	for i, num := range nums {
		if leftSum == sum-num-leftSum {
			return i
		}

		leftSum += num
	}

	return -1
}
