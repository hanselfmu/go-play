package algorithm

// ConvertZigZag solves LeetCode #6
func ConvertZigZag(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	strLen := len(s)
	r := make([]byte, strLen)
	charsPerRow := make([]int, numRows)
	insertIdxPerRow := make([]int, numRows)
	charsPerGroup := numRows*2 - 2

	mod, numOfGroups := strLen%charsPerGroup, strLen/charsPerGroup

	// numOfGroups does not count the last incomplete group
	if mod > 0 {
		charsPerRow[0] = numOfGroups + 1
	} else {
		charsPerRow[0] = numOfGroups
	}
	insertIdxPerRow[0] = 0
	insertIdxPerRow[1] = charsPerRow[0]
	for i := 1; i < numRows-1; i++ {
		if mod > i {
			charsPerRow[i] = 2*numOfGroups + 1
		} else {
			charsPerRow[i] = 2 * numOfGroups
		}
		// numRows: 4, mod: 5, i: 1, n
		// numRows: 4, mod: 5, i: 2, y
		// numRows: 5, mod: 7, i: 2, y
		// numRows: 5, mod: 7, i: 3, y
		// numRows: 5, mod: 7, i: 1, n
		// numRows: 6, mod: 8, i: 2, n
		// numRows: 6, mod: 8, i: 3, y
		if i+1+(mod-numRows) >= numRows {
			charsPerRow[i] = charsPerRow[i] + 1
		}
		insertIdxPerRow[i+1] = charsPerRow[i] + insertIdxPerRow[i]
	}

	// mod: 4, numRows: 4
	if mod >= numRows {
		charsPerRow[numRows-1] = numOfGroups + 1
	} else {
		charsPerRow[numRows-1] = numOfGroups
	}

	groupIdx := -1
	for idx, b := range []byte(s) {
		groupPos := idx % charsPerGroup
		if groupPos == 0 {
			groupIdx++
		}
		var rowIdx int
		if groupPos < numRows {
			rowIdx = groupPos
		} else {
			rowIdx = numRows - 2 - (groupPos - numRows)
		}

		curInsertIdx := insertIdxPerRow[rowIdx]
		r[curInsertIdx] = b
		insertIdxPerRow[rowIdx] = curInsertIdx + 1
	}

	return string(r)
}
