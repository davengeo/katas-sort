package impl

import "fmt"

type BinaryTree struct {
	value string
	left, right *BinaryTree
}

func newNode(k string) *BinaryTree {
	result := new(BinaryTree)
	result.value=k
	return result
}

func (me *BinaryTree) insertRight(node *BinaryTree) {
	me.right = node
}

func (me *BinaryTree) insertLeft(node *BinaryTree) {
	me.left = node
}

func (me *BinaryTree) greater(node *BinaryTree) (int) {
	if (me.value > node.value) {
		return 1
	}
	if (me.value < node.value) {
		return -1
	}
	return 0
}

func (me *BinaryTree) place(node *BinaryTree) {
	if (me.greater(node) == 1 && me.left == nil) {
		me.insertLeft(node)
		return
	}
	if (me.greater(node) == 1) {
		me.left.place(node)
		return
	}
	if (me.greater(node) == -1 && me.right == nil) {
		me.insertRight(node)
		return
	}
	if (me.greater(node) == -1) {
		me.right.place(node)
		return
	}
	if (me.greater(node) == 0) {
		if (me.right == nil) {
			me.insertRight(node)
			return
		}
		if (me.left == nil) {
			me.insertLeft(node)
			return
		}
		me.right.place(node)
	}
}

func findExtremeRight(me *BinaryTree, parent *BinaryTree) (*BinaryTree, *BinaryTree) {
	if (me.right == nil) {
		return me, parent
	}
	return findExtremeRight(me.right, me)
}

func flatRight(me *BinaryTree, arr []string) ([]string) {
	extreme, parent := findExtremeRight(me, nil)
	if (parent == nil) {
		return append(arr, extreme.value)
	}
	parent.right = extreme.left
	return flatRight(me, append(arr, extreme.value))
}

func findExtremeLeft(me *BinaryTree, parent *BinaryTree) (*BinaryTree, *BinaryTree) {
	if (me.left == nil) {
		return me, parent
	}
	return findExtremeLeft(me.left, me)
}

func flatLeft(me *BinaryTree, arr []string) ([]string) {
	extreme, parent := findExtremeLeft(me, nil)
	if (parent == nil) {
		return arr
	}
	parent.left = extreme.right
	return flatLeft(me, append(arr, extreme.value))
}

func reverse(arr []string) {
	for i, j := 0, len(arr) -1; i < j; i, j = i + 1, j -1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func sortTree(arr []string) ([]string, int)  {
	root := newNode(arr[0])
	for i := 1; i < len(arr); i++ {
		root.place(newNode(arr[i]))
	}
	right := flatRight(root, make([]string, 0))
	left := flatLeft(root, make([]string, 0))
	reverse(right)

	return append(left, right...), -1
}

func printArray(arr []string) {
	fmt.Println("*****")
	for i := 0; i< len(arr); i++ {
		fmt.Println(arr[i])
	}
}




