package average_salary_excluding_the_minimum_and_maximum_salary

func average(salary []int) float64 {
	sum := 0
	min := salary[0]
	max := salary[0]

	for _, s := range salary {
		if s < min {
			min = s
		} else if s > max {
			max = s
		}
		sum += s
	}

	return float64(sum-min-max) / float64(len(salary)-2)
}
