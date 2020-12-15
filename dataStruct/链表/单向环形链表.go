package main

import "fmt"

type catNode struct {
	no   int
	name string
	next *catNode
}

func addCatNode(head *catNode, newCat *catNode) {
	if head.next == nil {
		head.no = newCat.no
		head.name = newCat.name
		head.next = head
		return
	}

	temp := head

	for {
		if temp.next == head {
			newCat.next = head
			temp.next = newCat
			break
		}
		temp = temp.next
	}

}

func addCatNodeOrdered(head *catNode, newCat *catNode) {
	if head.next == nil {
		head.no = newCat.no
		head.name = newCat.name
		head.next = head
		return
	}
	temp := head

	for {
		if newCat.no == temp.next.no {
			fmt.Println("the cat you want to add existed")
			break
		}
		if newCat.no < temp.next.no {
			newCat.next = temp.next
			temp.next = newCat
			break
		}
		if temp.next == head {
			newCat.next = head
			temp.next = newCat
			break
		}
		temp = temp.next
	}

}

func deleteCatNode(head *catNode, id int) *catNode {
	if head.next == nil {
		fmt.Println("empty single ring link")
		return head
	}

	temp := head
	helper := head

	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	for {
		if id != head.no {
			if id == temp.next.no {
				temp.next = temp.next.next
				return head
			}
			if temp.next == head {
				fmt.Println("can't find node to delete")
				return head
			}
		}
		if id == head.no {
			helper.next = helper.next.next
			head = helper.next
			return head
		}
		temp = temp.next
	}

}

func displayCatNode(head *catNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("empty singleRingLink")
	}

	for {
		fmt.Printf("cat[%d][%s]-->", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	head := &catNode{}
	cat1 := &catNode{
		no:   1,
		name: "tom",
	}
	cat2 := &catNode{
		no:   2,
		name: "jack",
	}
	cat3 := &catNode{
		no:   3,
		name: "john",
	}
	cat4 := &catNode{
		no:   4,
		name: "anni",
	}

	addCatNode(head, cat1)
	addCatNode(head, cat2)
	addCatNode(head, cat3)
	addCatNode(head, cat4)
	displayCatNode(head)

	head = deleteCatNode(head, 4)
	displayCatNode(head)

}
