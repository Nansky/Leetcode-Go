// Given the root of a binary tree, return the inorder traversal of its nodes' values.

// Example 1:
// Input: root = [1,null,2,3]
// Output: [1,3,2]

// Example 2:
// Input: root = []
// Output: []

// Example 3:
// Input: root = [1]
// Output: [1]

// Constraints:
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	tempStack := make([]*TreeNode, 0)

	tempRoot := root

	for tempRoot != nil || len(tempStack) != 0 {
		for tempRoot != nil {
			tempStack = append(tempStack, tempRoot)
			tempRoot = tempRoot.Left
		}
		tempRoot = tempStack[len(tempStack)-1]
		res = append(res, tempRoot.Val)
		fmt.Println(tempRoot.Val)
		tempStack = tempStack[:len(tempStack)-1]
		tempRoot = tempRoot.Right
	}

	return res
}