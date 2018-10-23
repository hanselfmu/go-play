package algorithm

// TreeNode represents a BinaryTree Node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FixBST recovers a binary search tree that has two of its elements swapped.
// LeetCode #99
func FixBST(root *TreeNode) {
	inorder := traverseInorder(root)
	total := len(inorder)
	if total < 2 {
		return
	}

	var prev *TreeNode
	var prevIdx int
	anomalyIdx := -1
	for i, node := range inorder {
		if anomalyIdx == -1 {
			if prev != nil && node != nil {
				if prev.Val > node.Val {
					anomalyIdx = prevIdx
				}
			}
		}
		if anomalyIdx != -1 {
			// found the anomaly point; try changing with this element
			if i+1 == total {
				// last element; has to change
				temp := inorder[prevIdx].Val
				inorder[prevIdx].Val = inorder[anomalyIdx].Val
				inorder[anomalyIdx].Val = temp
			} else {
				if inorder[i+1] != nil && inorder[i+1].Val >= inorder[anomalyIdx].Val {
					// can switch
					// fmt.Printf("switching %d with %d\n", prevIdx, anomalyIdx)
					temp := inorder[prevIdx].Val
					inorder[prevIdx].Val = inorder[anomalyIdx].Val
					inorder[anomalyIdx].Val = temp
					return
				}
			}
		}
		if node != nil {
			prev = node
			prevIdx = i
		}
	}
}

func traverseInorder(node *TreeNode) []*TreeNode {
	if node == nil {
		return []*TreeNode{nil}
	}
	res := traverseInorder(node.Left)
	res = append(res, node)
	res = append(res, traverseInorder(node.Right)...)
	return res
}
