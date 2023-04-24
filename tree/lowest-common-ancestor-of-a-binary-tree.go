package tree

//236. 二叉树的最近公共祖先  && 剑指 Offer 68 - II. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

//235. 二叉搜索树的最近公共祖先(https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/)
//剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
func lowestCommonAncestorBinarySearchTree(root, p, q *TreeNode) *TreeNode {
	if p.Val < q.Val {
		if root == nil {
			return nil
		}

		if root == p || root == q {
			return root
		}

		if p.Val < root.Val && q.Val > root.Val {
			left := lowestCommonAncestor(root.Left, p, q)
			right := lowestCommonAncestor(root.Right, p, q)

			if left != nil && right != nil {
				return root
			}

			if left != nil {
				return left
			}
			return right
		} else if p.Val > root.Val {
			return lowestCommonAncestor(root.Right, p, q)
		} else {
			return lowestCommonAncestor(root.Left, p, q)
		}
	} else {
		return lowestCommonAncestorBinarySearchTree(root, q, p)
	}
}
