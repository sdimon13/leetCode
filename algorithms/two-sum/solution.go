package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, ok := m[complement]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}
