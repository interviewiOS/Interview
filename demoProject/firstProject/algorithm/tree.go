package algorithm

import (
	"fmt"
)

type TreeNode struct {
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

func CreateTrees() TreeNode {
	rootNode := TreeNode{
		1,
		nil,
		nil,
	}

	leftNode := TreeNode{
		2,
		nil,
		&TreeNode{4, nil, nil},
	}

	rightNode := TreeNode{
		3,
		&TreeNode{
			5,
			nil,
			&TreeNode{
				7, nil, nil,
			}},
		&TreeNode{6, &TreeNode{
			8, nil, nil,
		}, nil},
	}

	rootNode.leftNode = &leftNode
	rootNode.rightNode = &rightNode

	return rootNode
}

func Preoder(root *TreeNode) {
	if root == nil {
		fmt.Printf("root 为空", nil)
		return
	}
	fmt.Printf("%d ", root.value)
	if root.leftNode != nil {
		Preoder(root.leftNode)
	}
	if root.rightNode != nil {
		Preoder(root.rightNode)
	}
}

func Postorder(root *TreeNode) {
	if root == nil {
		fmt.Printf("root 为空", nil)
		return
	}
	if root.leftNode != nil {
		Postorder(root.leftNode)
	}
	if root.rightNode != nil {
		Postorder(root.rightNode)
	}
	fmt.Printf("%d ", root.value)
}

func Midorder(root *TreeNode) {
	if root == nil {
		fmt.Printf("root 为空", nil)
		return
	}
	if root.leftNode != nil {
		Midorder(root.leftNode)
	}
	fmt.Printf("%d ", root.value)
	if root.rightNode != nil {
		Midorder(root.rightNode)
	}

}
func LevelOrder(root *TreeNode) {
	if root == nil {
		return
	}
	lq := LinkQueue{
		nil,
		nil,
	}
	lList := []TreeNode{
		*root,
	}
	lq.addNode(lList)
	var targatCom = 2
	var targatDepth = 3
	var level = 0
	var targatNum = 0
	for !lq.isEmpty() {
		rootList, error := lq.outNode()
		if error != nil {
			continue
		}
		level++
		var iList []TreeNode
		for index, rootItem := range rootList.([]TreeNode) {
			fmt.Printf("%d ", rootItem.value)
			if index+1 == targatCom && targatDepth == level {
				targatNum = rootItem.value
			}
			if rootItem.leftNode != nil {
				iList = append(iList, *rootItem.leftNode)
			}
			if rootItem.rightNode != nil {
				iList = append(iList, *rootItem.rightNode)
			}
		}
		if len(iList) > 0 {
			lq.addNode(iList)
		}
	}
	fmt.Printf("\n深度：%d", level)
	fmt.Printf("\n第%d行，第%d个元素为%d ", targatDepth, targatCom, targatNum)
}

func LevelOrder2(root *TreeNode) {
	if root == nil {
		return
	}
	lq := LinkQueue{
		nil,
		nil,
	}

	lq.addNode(root)
	var targatCom = 2
	var targatDepth = 2
	var level = 1
	var targatNum = 0
	var lENode = root
	var index = 0
	for !lq.isEmpty() {
		var lSNode *TreeNode
		rootItem, error := lq.outNode()
		if error != nil {
			continue
		}
		iNode := rootItem.(*TreeNode)
		fmt.Printf("%d ", iNode.value)
		if iNode.leftNode != nil {
			lSNode = iNode.leftNode
			lq.addNode(iNode.leftNode)
		}
		if iNode.rightNode != nil {
			lSNode = iNode.rightNode
			lq.addNode(iNode.rightNode)
		}

		if level == targatDepth && index+1 == targatCom {
			targatNum = iNode.value
		}

		if rootItem == lENode {
			level++
			lENode = lSNode
			index = 0
		} else {
			index++
		}

	}
	fmt.Printf("\n深度：%d", level)
	fmt.Printf("\n第%d行，第%d个元素为%d ", targatDepth, targatCom, targatNum)
}

func TssTree() {
	node := CreateTrees()
	print("前序 ")
	Preoder(&node)
	print("\n中序 ")
	Midorder(&node)
	print("\n后序 ")
	Postorder(&node)
	print("\n层级 ")
	//LevelOrder(&node)
	LevelOrder2(&node)
}
