/*

This was made by Antons but it doesnt work, his email is akranga@gmail.com

 */
package impl

import "fmt"

type BinTree struct {
	value Any
	left, right, parent *BinTree
}

func (first BinTree) compareTree (second BinTree) int {
	strFirst := first.value.toString()
	strSecond := second.value.toString()

	if (strFirst > strSecond) {
		return 1
	}
	if (strFirst < strSecond){
		return -1
	}
	return 0
}


func newBinTree(value Any) BinTree {
	node := new(BinTree)
	node.value = value
	node.parent = nil
	node.left = nil
	node.right = nil
	return *node
}

func (curr BinTree) place(node BinTree) {
	result := curr.compareTree(node)
	if result < 0 {
		if (curr.left == nil) {
			curr.left = &node
			node.parent = &curr
		} else {
			curr.left.place(node)
		}
	} else {
		if (curr.right == nil) {
			curr.right = &node
			node.parent = &curr
		} else {
			curr.right.place(node)
		}
	}
}

func scanLeft(node *BinTree) (BinTree) {
	if (node.left == nil) {
		return *node
	}
	return scanLeft(node.left)
}

func TreeSort(arr []Any) ([]Any, int) {
	root := newBinTree(arr[0])
	curr := root
	//for i := 1; i < (len(arr) - 1); i++ {
	for i := 1; i < 10; i++ {
		other := newBinTree(arr[i+1])
		curr.place(other)
		fmt.Println(curr)
	}

	var result = make([]Any, len(arr))
	left := scanLeft(&root)
	fmt.Println("****")
	fmt.Println(left)
	leftToRightCopy(left, result, 0)
	return result, -1
}

func leftToRightCopy(curr BinTree, result []Any, i int) {
	if (curr.left != nil) {
		leftToRightCopy(*(curr.left), result, i)
	}
	result[i] = curr.value

	i++
	if (curr.right != nil) {
		leftToRightCopy(*(curr.right), result, i)
	} else if (curr.parent != nil) {
		leftToRightCopy(*(curr.parent), result, i)
	}
}