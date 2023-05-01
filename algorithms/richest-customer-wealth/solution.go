package richest_customer_wealth

import "sync"

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, acc := range accounts {
		sum := 0
		for _, money := range acc {
			sum += money
		}

		if sum > max {
			max = sum
		}
	}
	return max
}

func maximumWealthV2(accounts [][]int) int {
	maxSum := 0
	ch := make(chan int, len(accounts))
	wg := sync.WaitGroup{}
	wg.Add(len(accounts))

	for _, acc := range accounts {
		go calculateWealth(acc, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for accountsum := range ch {
		if accountsum > maxSum {
			maxSum = accountsum
		}
	}

	return maxSum
}

func calculateWealth(account []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, money := range account {
		sum += money
	}

	ch <- sum
}
