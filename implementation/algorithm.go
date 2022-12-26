package main

import (
	"fmt"

	// "github.com/pxsa/iGraph"
)

type piziStar struct {
	states   []myNode
	openList []myNode
	start    myNode
	goal     myNode
}

type myNode struct {
	tag        string
	// node       iGraph.Node
	neighbors  []*myNode
	next       *myNode
	costToGoal int
}

func (pizi *piziStar) process_state() int {
	x := pizi.min_state()
	if x == nil {
		return -1
	}

	k_old := pizi.get_kmin()
	delete(x)

	if k_old < x.h() {
		for _, neighbor := range x.neighbors {
			if neighbor.h() <= k_old && x.h() > neighbor.h() + pizi.c(neighbor, x) {
				x.next = neighbor
				x.costToGoal = neighbor.h() + pizi.c(neighbor, x)
			}
		}
	} else if k_old == x.h() {
		fmt.Println("")
	} else {
		fmt.Println("")
	}

	return pizi.get_kmin()
}

// returns the state on the OPENLIST with minimum k(.) value
// (NULL if the list is empty)
func (pizi *piziStar) min_state() *myNode {
	// for node := range pizi.openList {
	// }
	return &pizi.openList[0]
}

// returns Kmin for the OPENLIST (-1 if the list is empty)
func (pizi *piziStar) get_kmin() int {

	return pizi.min_state().k()
}

func delete(node *myNode) {

}

func (node *myNode) h() int {
	return node.costToGoal
}

// func (node *myNode) b() *myNode {
// 	return node.next
// }

func (pizi *piziStar) c(a, b *myNode) int {
	var val int
	return val
}

func (node *myNode) k() int {
	return 1
}
