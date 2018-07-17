package algorithm

// RepeatToContain calculates how many times "a" needs to be repeated
// to be able to contain "b" as a substring
// from https://www.geeksforgeeks.org/google-interview-experience-set-7-software-engineering-intern/
func RepeatToContain(a, b string) int {
	seqList := findIndex(a, b[0])
	timesList := make([]int, len(seqList))
	for i := range timesList {
		timesList[i] = 1
	}
	lenA := len(a)

	if len(seqList) == 0 {
		return -1
	}

	bytesB := []byte(b)
	for _, vB := range bytesB[1:] {
		if len(seqList) == 0 {
			return -1
		}

		newSeqList := make([]int, 0, len(seqList))
		newTimesList := make([]int, 0, len(seqList))
		for i, curSeq := range seqList {
			newSeq := curSeq + 1
			curTimes := timesList[i]
			if newSeq >= lenA {
				newSeq = 0
				curTimes = curTimes + 1
			}
			if a[newSeq] == vB {
				newSeqList = append(newSeqList, newSeq)
				newTimesList = append(newTimesList, curTimes)
			}
		}

		seqList = newSeqList
		timesList = newTimesList
	}

	res := -1
	for _, v := range timesList {
		if v < res || res < 0 {
			res = v
		}
	}
	return res
}

func findIndex(s string, b byte) []int {
	res := make([]int, 0, len(s))
	for i, v := range []byte(s) {
		if b == v {
			res = append(res, i)
		}
	}
	return res
}
