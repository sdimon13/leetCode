package bulls_and_cows

import "fmt"

func getHint(secret string, guess string) string {
	digits := [10]int{}
	bulls, cows := 0, 0

	for i := 0; i < len(secret); i++ {
		s := secret[i] - '0'
		g := guess[i] - '0'

		if s == g {
			bulls++
		} else {
			if digits[s] < 0 {
				cows++
			}
			if digits[g] > 0 {
				cows++
			}

			digits[s]++
			digits[g]--
		}

	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
