package algorithm

// LemonadeChange solves #860 of LeetCode algorithm challenges
func LemonadeChange(bills []int) bool {
	cash5 := 0
	cash10 := 0

	for _, bill := range bills {
		switch bill {
		case 5:
			cash5++
		case 10:
			cash10++
			if cash5 <= 0 {
				return false
			}
			cash5--
		case 20:
			if cash5 <= 0 {
				return false
			}
			if cash10 <= 0 {
				if cash5 < 3 {
					return false
				}
				cash5 -= 3
			} else {
				if cash5 < 1 {
					return false
				}
				cash10--
				cash5--
			}

		}
	}

	return true
}
