package main

import "fmt"

type GOTnode2 struct {
	no   int
	name string
	prev *GOTnode2
	next *GOTnode2
}

func addNode2(head *GOTnode2, newNode *GOTnode2) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	temp.next = newNode
	newNode.prev = temp

}

func addNodeOrdered2(head *GOTnode2, newNode *GOTnode2) {
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
		newNode.prev = temp
		if temp.next != nil {
			temp.next.prev = newNode
		}
		temp.next = newNode
	}
}

func deleteNode2(head *GOTnode2, id int) {
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
		if temp.next != nil {
			temp.next.prev = temp
		}
	} else {
		fmt.Println("the node you search for doesn't exist")
	}

}

func displayNode2(head *GOTnode2) {
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

func displayNode2R(head *GOTnode2) {
	temp := head
	if temp.next == nil {
		fmt.Println("empty double list")
		return
	}

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	for {
		if temp.prev == nil {
			fmt.Println("empty space")
			break
		}
		fmt.Printf("[%d,%s]-->", temp.no, temp.name)
		temp = temp.prev
	}
	fmt.Println()
}

func main() {
	head := &GOTnode2{}
	node1 := &GOTnode2{
		no:   1,
		name: "arya",
	}
	node2 := &GOTnode2{
		no:   2,
		name: "jon",
	}
	node3 := &GOTnode2{
		no:   3,
		name: "sansa",
	}

	/*	addNode2(head, node1)
		addNode2(head, node2)
		addNode2(head, node3)*/

	addNodeOrdered2(head, node3)
	addNodeOrdered2(head, node1)
	addNodeOrdered2(head, node2)

	displayNode2(head)

	displayNode2R(head)

	deleteNode2(head, 2)
	displayNode2(head)

}
