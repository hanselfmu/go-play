package fundamental

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
		bt.Insert(val)
	}

	return bt
}

// Insert inserts a new value into a binary tree, *iteratively*
func (bt *BinaryTree) Insert(val Comparable) {
	for inserted := false; !inserted; {
		if bt == nil {
			bt = &BinaryTree{val, nil, nil}
			inserted = true
		} else {
			if bt.Value.LessThan(val) {
				bt = bt.Right
			} else {
				bt = bt.Left
			}
		}
	}
}
