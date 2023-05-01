package number_of_steps_to_reduce_a_number_to_zero

func numberOfSteps(num int) int {
	i := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		i++
	}
	return i
}

func numberOfStepsV2(num int) int {
	if num < 2 {
		return num
	}
	i := 1
	for num > 1 {
		i += 1 + (num & 1)
		num >>= 1
	}
	return i
}
