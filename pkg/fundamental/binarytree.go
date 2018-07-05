package fundamental

import "fmt"

// BinaryTree is a binary tree node
type BinaryTree struct {
	Value       Comparable
	Left, Right *BinaryTree
}

// NewBinaryTree creates a new binary tree
func NewBinaryTree(vals []Comparable) *BinaryTree {
	return new(BinaryTree).Init(vals)
}

// Init initializes a new binary tree with an array of data
func (bt *BinaryTree) Init(vals []Comparable) *BinaryTree {
	for _, val := range vals {
		fmt.Printf("inserting %d for tree %v at %p\n", val, *bt, bt)
		bt.Insert(val)
	}

	return bt
}

// Insert inserts a new value into a binary tree, *iteratively*
func (bt *BinaryTree) Insert(val Comparable) {
	for inserted := false; !inserted; {
		if bt.Value == nil {
			bt.Value = val
			inserted = true
		} else {
			if bt.Value.LessThan(val) {
				if bt.Right == nil {
					bt.Right = &BinaryTree{nil, nil, nil}
				}

				bt = bt.Right
			} else {
				if bt.Left == nil {
					bt.Left = &BinaryTree{nil, nil, nil}
				}

				bt = bt.Left
			}
		}
	}
}

// DepthRec returns the depth of a binary tree, recursively
func (bt *BinaryTree) DepthRec() int {
	if bt == nil {
		return 0
	}

	return Max(1+bt.Left.DepthRec(), 1+bt.Right.DepthRec())
}

// DepthIte returns the depth of a binary tree, iteratively
func (bt *BinaryTree) DepthIte() int {
	depth := 0
	curLevel := []*BinaryTree{bt}
	newLevel := make([]*BinaryTree, 0, 2)
	idx := 0
	curLen := len(curLevel)

	for curLen > 0 {
		node := curLevel[idx]
		if node.Left != nil {
			newLevel = append(newLevel, node.Left)
		}
		if node.Right != nil {
			newLevel = append(newLevel, node.Right)
		}

		idx++
		if idx == curLen {
			// end of current level; switch
			depth++
			curLevel = newLevel
			newLevel = make([]*BinaryTree, 0, curLen*2)
			idx = 0
			curLen = len(curLevel)
		}
	}

	return depth
}

// Beautify formats binary tree and generates a formatted string
// func (bt *BinaryTree) Beautify() string {

// }
