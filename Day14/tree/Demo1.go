package main

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
	"sync"
)

/*

			----[ 1
		----[ 2
			----[ 3
	----[ 4
			----[ 5
		----[ 6
			----[ 7
----[ 8
		----[ 9
	----[ 10
		----[ 11
*/

// Item interface{}
type Item generic.Type

// 节点结构体
type Node struct {
	key   int   //key,根据key进行节点插入
	value Item  //Item interface{}
	left  *Node //left
	right *Node //right
}

// 二叉树结构体
type ItemBinarySearchTree struct {
	root *Node        // 根节点
	lock sync.RWMutex //读写锁
}

// 格式化输出二叉搜索树
func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// 通过递归进行格式化输出
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}

// 构建ItemBinarySearchTree搜索二叉树插入节点的方法
// 有两个参数，一个是节点的key，另一个是value
func (tree *ItemBinarySearchTree) Insert(key int, value Item) {

	// 上锁
	tree.lock.Lock()
	// 解锁
	defer tree.lock.Unlock()

	// 创建一个节点
	node := &Node{key, value, nil, nil}

	if tree.root == nil {
		tree.root = node
	} else {
		// 以根节点为入口，插入新节点node
		insertNode(tree.root, node)
	}
}

func insertNode(node, newNode *Node) {
	// 判断节点插入到左边还是右边
	// 如果新节点的key小于根节点的key，插入到左边，否则插入到右边
	if newNode.key < node.key {

		if node.left == nil {
			node.left = newNode
		} else {
			// 通过递归插入新节点到左边
			insertNode(node.left, newNode)
		}

	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			// 通过递归插入新节点到右边
			insertNode(node.right, newNode)
		}

	}
}

func main() {
	// 创建二叉搜索树对象
	// 0xc4200a4020
	var tree ItemBinarySearchTree
	// 打印tree
	fmt.Println(tree)
	fmt.Printf("%p",&tree)
	// 插入第一个节点
	//tree.Insert(8, "8")
	//tree.Insert(4, "4")
	//tree.Insert(10, "10")
	//tree.Insert(2, "2")
	//tree.Insert(6, "6")
	//tree.Insert(1, "1")
	//tree.Insert(3, "3")
	//tree.Insert(5, "5")
	//tree.Insert(7, "7")
	//tree.Insert(9, "9")
	//tree.Insert(11, "11")

	tree.String()

}
