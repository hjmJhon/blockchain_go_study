package main

import "fmt"

//定义一个节点
type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

//定义一个根节点
type RootNode struct {
	node *Node
}

//插入到节点
func (tree *RootNode)insertNode(key,value int) {

	var node = &Node{key:key,value:value}
	if tree.node == nil {
		tree.node = node
		fmt.Println()
	} else {
		insertSubNode(node,tree.node)
	}

}

//插入到子节点
func insertSubNode(newNode *Node,fatherNode *Node) {
	if newNode.key > fatherNode.key {
		fmt.Print("\t\t")
		//新节点插入到父节点的右边
		if fatherNode.right ==nil {
			fatherNode.right = newNode
		}else {
			insertSubNode(newNode,fatherNode.right)
		}

	}else {
		//新节点插入到父节点左边
		if fatherNode.left ==nil {
			fatherNode.left = newNode
		}else {
			insertSubNode(newNode,fatherNode.left)
		}
	}
}

func main() {
	tree := &RootNode{}
	fmt.Println(tree)
	tree.insertNode(8,8)

}
