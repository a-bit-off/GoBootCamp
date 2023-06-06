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
	HasToy: false,
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
	HasToy: false,
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

// 	     0
//	   /   \
//		0     1
//	 /\	    /\
//  0  1	 0  1

var head3 = TreeNode{
	HasToy: false,
	Left: &TreeNode{
		HasToy: false,
		Left:   &TreeNode{HasToy: false, Left: nil, Right: nil},
		Right:  &TreeNode{HasToy: true, Left: nil, Right: nil},
	},
	Right: &TreeNode{
		HasToy: true,
		Left:   &TreeNode{HasToy: false, Left: nil, Right: nil},
		Right:  &TreeNode{HasToy: true, Left: nil, Right: nil},
	},
}

func main() {
	print(unrollGarland(&head2))
}

func print(result [][]bool) {
	for _, res := range result {
		for _, r := range res {
			fmt.Printf("%t ", r)
		}
		fmt.Println()
	}
}

func snakeVersion(result [][]bool) {
	for i := 0; i < len(result); i++ {
		if i%2 == 0 {
			ln := len(result[i]) - 1
			for j := ln; j > ln/2; j-- {
				curr := result[i][ln-j]
				result[i][ln-j] = result[i][j]
				result[i][j] = curr
			}
		}
	}
}

func unrollGarland(head *TreeNode) [][]bool {
	if head == nil {
		return [][]bool{}
	}
	result := [][]bool{}
	queue := []*TreeNode{head}
	for len(queue) > 0 {
		generation := len(queue)
		childrenGeneration := []bool{}

		for i := 0; i < generation; i++ {
			node := queue[0]
			queue = queue[1:]
			childrenGeneration = append(childrenGeneration, node.HasToy)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, childrenGeneration)
	}
	snakeVersion(result)
	return result
}
