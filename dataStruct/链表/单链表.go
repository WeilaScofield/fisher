package main

import "fmt"

type GOTnode struct {
	no   int
	name string
	next *GOTnode
}

func addNode(head *GOTnode, newNode *GOTnode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	temp.next = newNode
}

func addNodeOrdered(head *GOTnode, newNode *GOTnode) {
	temp := head
	flag := true

	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newNode.no {
			break
		} else if temp.next.no == newNode.no {
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("node already exist")
	} else {
		newNode.next = temp.next
		temp.next = newNode
	}
}

func deleteNode(head *GOTnode, id int) {
	temp := head
	flag := false

	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}

	// GC will collect the garbage node
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("the node you search for doesn't exist")
	}

}

func displayNode(head *GOTnode) {
	temp := head

	for {
		if temp.next == nil {
			fmt.Println("empty space")
			break
		}
		fmt.Printf("[%d,%s]-->", temp.next.no, temp.next.name)
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	head := &GOTnode{}
	node1 := &GOTnode{
		no:   1,
		name: "arya",
	}
	node2 := &GOTnode{
		no:   2,
		name: "jon",
	}
	node3 := &GOTnode{
		no:   3,
		name: "sansa",
	}

	addNodeOrdered(head, node3)
	addNodeOrdered(head, node1)
	addNodeOrdered(head, node2)

	displayNode(head)

	deleteNode(head, 2)
	displayNode(head)
}
