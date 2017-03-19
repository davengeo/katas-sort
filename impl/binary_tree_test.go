package impl

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateFreshBinaryTree(t *testing.T) {
	root := newNode("Deniss")
	assert.Nil(t, root.left)
	assert.Nil(t, root.right)
	assert.Equal(t, "Deniss", root.value)

	fidan:=newNode("Fidan")

	root.insertRight(fidan)
	assert.Nil(t,root.left)
	assert.Equal(t, root.right.value, fidan.value)

	mathieu:=newNode("Mathieu")
	fidan.insertRight(mathieu)
	assert.Equal(t, root.right.right.value, mathieu.value)

	assert.Equal(t, fidan.greater(root), 1)

	david:=newNode("David")
	root.place(david)

	assert.Equal(t, root.left.value, david.value)

	florian:=newNode("Florian")
	root.place(florian)
	assert.Equal(t, root.right.right.left.value, florian.value)
	extreme, parent := findExtremeRight(root, nil)
	assert.Equal(t, extreme.value, mathieu.value)
	assert.Equal(t, parent.value, fidan.value)

	right := flatRight(root, make([]string, 0))
	assert.Equal(t, right[0], "Mathieu")
	assert.Equal(t, right[1], "Florian")
	assert.Equal(t, right[2], "Fidan")
	assert.Equal(t, right[3], "Deniss")

	left := flatLeft(root, make([]string, 0))
	assert.Equal(t, left[0], "David")
	assert.Equal(t, len(left), 1)

	reverse(right)
	printArray(append(left, right...))

	arr, _ := SortTree([]string{"Zorba", "David", "Olli", "Evegenijs", "David", "Antons", "Ivo", "Deniss"})
	printArray(arr)
}