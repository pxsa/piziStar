package main

type piziNode struct {
	row        int
	col        int
	tag        string
	costToGoal int // h
	key        int // k
	isObstacle bool
	next       *piziNode // b
	neighbors  []*piziNode
}

func (node *piziNode) h() int {
	return node.costToGoal
}

func (node *piziNode) b() *piziNode {
	return node.next
}

func (node *piziNode) k() int {
	return node.key
}
