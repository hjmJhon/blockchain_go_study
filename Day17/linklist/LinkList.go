package linklist

import (
	"fmt"
)

type Node struct {
	//节点数据
	data interface{}
	//引用的下一个节点
	next *Node
}

type LinkList struct {
	size int
	head *Node
	tail *Node
}

func CreateLinkList() *LinkList {
	return &LinkList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

/*
	追加 数据到链表尾部
 */
func (linkList *LinkList) Append(data interface{}) bool {
	node := &Node{
		data: data,
		next: nil,
	}

	if linkList.size == 0 {
		linkList.head = node
		linkList.tail = node
	} else {
		oldTail := linkList.tail
		oldTail.next = node
		oldTail = node
	}
	linkList.size++

	return true

}

/*
	插入到指定位置
 */
func (linkList *LinkList) Insert(index int, data interface{}) error {
	if linkList.size < index || linkList.size == 0 || index < 0 {
		return fmt.Errorf("index error or out of the LinkList's size");
	}

	newNode := &Node{
		data: data,
		next: nil,
	}
	if index == 0 {
		newNode.next = linkList.head
		linkList.head = newNode
	} else {
		node := linkList.head
		for i := 1; i < index; i++ {
			node = node.next
		}
		newNode.next = node.next
		node.next = newNode
	}

	linkList.size++

	return nil
}

/*
	获取指定的节点
 */
func (linkList *LinkList) IndexOf(index int) *Node {
	if index > linkList.size {
		return nil
	}

	node := linkList.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node

}

func (linkList *LinkList) Head() *Node {
	return linkList.head
}

func (linkList *LinkList) Tail() *Node {
	return linkList.tail
}

func (linkList *LinkList) Remove(index int) bool {
	if linkList.size == 0 || index < 0 || index >= linkList.size {
		return false
	}
	if index == 0 {
		linkList.head = linkList.head.next
	} else {
		node := linkList.head
		for i := 1; i < index; i++ {
			node = node.next
		}
		node.next = node.next.next
	}

	linkList.size--

	return true
}

func (linkList *LinkList) Size() int {
	return linkList.size
}

func (linkList *LinkList) String() {
	fmt.Println("LinkList String:")
	head := linkList.head
	for head != nil {
		fmt.Print(head.data)
		head = head.next
		if head != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}
