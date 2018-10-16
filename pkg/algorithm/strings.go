package algorithm

// LongestValidParentheses finds the length of the longest valid (well-formed) parentheses substring.
// The parameter "s" contains just the characters '(' and ')'.
// LeetCode #32
func LongestValidParentheses(s string) int {
	globalMax := -1
	var stackTop *listNode
	var stackEmptyNodeTop *emptyNode
	hasCalculatedCurrentRun := false
	for _, char := range s {
		if char == '(' {
			hasCalculatedCurrentRun = false
			// push onto stack a new "("
			newNode := &listNode{false, stackTop}
			newEmptyNode := &emptyNode{newNode, stackEmptyNodeTop}
			stackTop = newNode
			stackEmptyNodeTop = newEmptyNode
		} else {
			if stackEmptyNodeTop != nil {
				// has empty node to pop to match current ")"
				stackEmptyNodeTop.lNode.occupied = true
				stackEmptyNodeTop = stackEmptyNodeTop.next
			} else {
				if !hasCalculatedCurrentRun {
					hasCalculatedCurrentRun = true
					curRunMax := 0
					curTally := 0
					curNode := stackTop
					for curNode != nil {
						if curNode.occupied {
							curTally += 2
						} else {
							if curTally > curRunMax {
								curRunMax = curTally
							}
							curTally = 0
						}
						curNode = curNode.next
					}
					if curTally > curRunMax {
						curRunMax = curTally
					}
					if curRunMax > globalMax {
						globalMax = curRunMax
					}

					stackTop = nil
					stackEmptyNodeTop = nil
				}
			}
		}
	}

	if !hasCalculatedCurrentRun {
		hasCalculatedCurrentRun = true
		curRunMax := 0
		curTally := 0
		curNode := stackTop
		for curNode != nil {
			if curNode.occupied {
				curTally += 2
			} else {
				if curTally > curRunMax {
					curRunMax = curTally
				}
				curTally = 0
			}
			curNode = curNode.next
		}
		if curTally > curRunMax {
			curRunMax = curTally
		}
		if curRunMax > globalMax {
			globalMax = curRunMax
		}
	}

	return globalMax
}

type listNode struct {
	occupied bool
	next     *listNode
}

type emptyNode struct {
	lNode *listNode
	next  *emptyNode
}
