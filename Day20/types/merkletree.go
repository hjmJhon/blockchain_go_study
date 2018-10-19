package types

import (
	"crypto/sha256"
	"log"
)

type MerkleNode struct {
	LeftNode  *MerkleNode
	RightNode *MerkleNode
	Hash      []byte
}

type MerkleTree struct {
	RootNode *MerkleNode
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	datas := data
	if len(datas)%2 != 0 {
		datas = append(datas, datas[len(datas)-1])
	}

	var nodes []*MerkleNode

	//创建叶子节点
	for _, d := range datas {
		node := newMerkleNode(nil, nil, d)
		nodes = append(nodes, node)
	}

	//创建非叶子节点
	if len(nodes)%2 != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
	}
	for i := 0; i < len(nodes)/2; i++ {
		var nodeArr []*MerkleNode
		for j := 0; j < len(nodes); j += 2 {
			node := newMerkleNode(nodes[j], nodes[j+1], nil)
			nodeArr = append(nodeArr, node)
		}

		if len(nodeArr)%2 != 0 {
			nodeArr = append(nodeArr, nodeArr[len(nodeArr)-1])
		}
		nodes = nodeArr
	}

	return &MerkleTree{RootNode: nodes[0]}
}

func newMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	if (left == nil && right != nil) || (left != nil && right == nil) {
		log.Panic("error: invalidate node")
	}

	var merkleNode = &MerkleNode{}
	if left == nil && right == nil {
		sum := sha256.Sum256(data)
		merkleNode.Hash = sum[:]
	} else {
		dataArr := append(left.Hash, right.Hash...)
		sum := sha256.Sum256(dataArr)
		merkleNode.Hash = sum[:]
	}

	merkleNode.LeftNode = left
	merkleNode.RightNode = right

	return merkleNode
}
