package main

type treeNode struct {
	data     int
	leftPtr  *treeNode
	rightPtr *treeNode
}

type binaryTree struct {
	root *treeNode
}

func createNode(val int) *treeNode {

	var node treeNode
	node.data = val
	node.leftPtr = nil
	node.rightPtr = nil

	return &node
}

func buildTree(arr []int, idx int, root *treeNode) *treeNode {

	if arr[idx] == -1 {
		return nil
	} else {
		return nil
	}

}

func main() {

	var data = []int{2, 3, 4, -1, 2, 5, -1}

	var root *treeNode = nil

	root = buildTree(data, 0, root)

}
