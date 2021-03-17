package main

import "fmt"

func RunBranchSums() {
	tree := newBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sums := branchSums(tree)
	fmt.Printf("%v\n", sums)
}

//from here until func branchSums() borrowed for easy test building
func newBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}
	tree.Insert(values, 0)
	return tree
}

func (tree *BinaryTree) Insert(values []int, i int) *BinaryTree {
	if i >= len(values) {
		return tree
	}
	val := values[i]

	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		var current *BinaryTree
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Right)
	}

	tree.Insert(values, i+1)
	return tree
}

//code below is mine. code above was used from sample test cases. review it.
func branchSums(root *BinaryTree) []int {
	sums := []int{}
	branchSumsRecursive(root, &sums, 0)
	return sums
}

func branchSumsRecursive(root *BinaryTree, sums *[]int, sum int) {
	if root == nil {
		return
	}
	sum += root.Value
	if root.Left == nil && root.Right == nil {
		*sums = append(*sums, sum)
		return
	}
	branchSumsRecursive(root.Left, sums, sum)
	branchSumsRecursive(root.Right, sums, sum)
}

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}
