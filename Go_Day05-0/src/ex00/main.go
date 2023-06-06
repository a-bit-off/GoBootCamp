package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

// 	    0
//	   / \
//		0   1
//	 / \
//  0   1

var head1 = TreeNode{
	HasToy: true,
	Left: &TreeNode{
		HasToy: false,
		Left:   &TreeNode{HasToy: false, Left: nil, Right: nil},
		Right:  &TreeNode{HasToy: true, Left: nil, Right: nil},
	},
	Right: &TreeNode{HasToy: true, Left: nil, Right: nil},
}

// 	    0
//	   / \
//		0   1
//	 / \	 \
//  0   1		1

var head2 = TreeNode{
	HasToy: true,
	Left: &TreeNode{
		HasToy: false,
		Left:   &TreeNode{HasToy: false, Left: nil, Right: nil},
		Right:  &TreeNode{HasToy: true, Left: nil, Right: nil},
	},
	Right: &TreeNode{
		HasToy: true,
		Left:   nil,
		Right:  &TreeNode{HasToy: true, Left: nil, Right: nil},
	},
}

func main() {
	fmt.Println("areToysBalanced:", areToysBalanced(&head1))
	fmt.Println("areToysBalanced:", areToysBalanced(&head2))
}

func areToysBalanced(head *TreeNode) bool {
	leftCount, rightCount := counter(head.Left), counter(head.Right)
	return leftCount == rightCount
}

func counter(head *TreeNode) int {
	count := 0
	if head != nil {
		if head.HasToy {
			count = 1
		}
		count += counter(head.Left)
		count += counter(head.Right)
	}
	return count
}
