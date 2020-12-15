package main

import (
	"fmt"
	"os"
)

type emp struct {
	id   int
	next *emp
}

type empHead struct {
	head *emp
}

type hashTab struct {
	arr [7]empHead
}

//在empHead下按照顺序添加雇员
func (e *empHead) insert(em *emp) {
	cur := e.head
	var pre *emp

	//如果empHead后面为空，直接添加雇员
	if cur == nil {
		e.head = em
		return
	}

	//如果empHead后面不为空
	for {
		if cur != nil {
			if cur.id > em.id { //找到了雇员应该插入的位置
				break
			}
			pre = cur //cur和pre同步往后走
			cur = cur.next
		} else {
			break //如果走到了链表的末端，则break
		}
	}

	//将雇员在合适位置插入
	em.next = cur
	pre.next = em
}

//显示当前链表
func (e *empHead) displayEmpHead(i int) {
	cur := e.head
	if cur == nil {
		fmt.Printf("arr[%d]:empty link\n", i)
		return
	}

	for {
		if cur != nil {
			fmt.Printf("arr[%d]:userId[%d]-->", i, cur.id)
			cur = cur.next
		} else {
			break
		}
	}
	fmt.Println()
}

//显示所有雇员
func (h *hashTab) displayAll() {
	for i := 0; i < len(h.arr); i++ {
		h.arr[i].displayEmpHead(i)
	}
}

//将雇员通过hash函数，散列到数组的不同位置
func (h *hashTab) insert(emp *emp) {
	index := h.hashFunc(emp.id)
	h.arr[index].insert(emp)
}

func (h *hashTab) hashFunc(id int) int {
	return id % 7
}

func (e *empHead) findEmp(id int) *emp {
	cur := e.head
	if cur == nil {
		fmt.Println("empty link")
		return nil
	}

	for {
		if cur.id == id {
			return cur
		}
		cur = cur.next
		if cur == nil {
			return nil
		}
	}
}

func (h *hashTab) findEmp(id int) *emp {
	index := h.hashFunc(id)
	e := h.arr[index].findEmp(id)
	return e
}

func main() {
	var id int
	var key int
	h := &hashTab{}
	for {
		fmt.Println("employee information system")
		fmt.Println("1:show all")
		fmt.Println("2:insert employee")
		fmt.Println("3:find employee")
		fmt.Println("4:exit")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			h.displayAll()
		case 2:
			fmt.Println("please enter userID")
			fmt.Scanln(&id)
			e := &emp{
				id: id,
			}
			h.insert(e)
		case 3:
			fmt.Println("please enter the userID you wanna find out")
			fmt.Scanln(&id)
			e := h.findEmp(id)
			if e == nil {
				fmt.Println("haven't found out employee")
			} else {
				fmt.Println(e.id, "has been found")
			}
		case 4:
			os.Exit(1)
		}
		fmt.Println()
	}

}
