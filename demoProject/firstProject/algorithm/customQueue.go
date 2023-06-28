package algorithm

import "errors"

type LinkNode struct {
	Value    interface{}
	NextNode *LinkNode
}

type LinkQueue struct {
	FrontNode *LinkNode
	LastNode  *LinkNode
}

func (lq *LinkQueue) addNode(value interface{}) {
	newNode := &LinkNode{
		value,
		nil,
	}
	if lq.FrontNode == nil {
		lq.FrontNode = newNode
		lq.LastNode = newNode
	} else {
		lq.LastNode.NextNode = newNode
		lq.LastNode = newNode
	}
}

func (lq *LinkQueue) outNode() (interface{}, error) {
	if lq.FrontNode == nil {
		return -1, errors.New("对列为空")
	}
	fNode := lq.FrontNode
	lq.FrontNode = lq.FrontNode.NextNode
	return fNode.Value, nil
}

func (lq *LinkQueue) isEmpty() bool {
	if lq.FrontNode == nil {
		return true
	}
	return false
}
