package main

import "fmt"

type tree struct {
	val   int
	left  *tree
	right *tree
}

func newTree(v int) *tree {
	binaryTree := &tree{val: v}
	return binaryTree
}

func (t *tree) preOrder() {
	if t != nil {
		fmt.Printf("%d\t", t.val)
		t.left.preOrder()
		t.right.preOrder()
	}
}

func (t *tree) midOrder() {
	if t != nil {
		t.left.midOrder()
		fmt.Printf("%d\t", t.val)
		t.right.midOrder()
	}
}

func (t *tree) postOrder() {
	if t != nil {
		t.left.postOrder()
		t.right.postOrder()
		fmt.Printf("%d\t", t.val)
	}
}

func main() {
	tree1 := newTree(1)
	tree2 := newTree(2)
	tree3 := newTree(3)
	tree4 := newTree(4)
	tree5 := newTree(5)
	tree6 := newTree(6)
	tree7 := newTree(7)

	tree1.left = tree2
	tree1.right = tree3
	tree2.left = tree4
	tree2.right = tree5
	tree3.left = tree6
	tree3.right = tree7

	fmt.Println("preOrder--")
	tree1.preOrder()
	fmt.Println()
	fmt.Println("midOrder--")
	tree1.midOrder()
	fmt.Println()
	fmt.Println("postOrder--")
	tree1.postOrder()
}
