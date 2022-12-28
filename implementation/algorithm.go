package main

import (
	"fmt"
	"math"

	"github.com/pxsa/iGraph"
	"golang.org/x/exp/slices"
)

type piziStar struct {
	states   []*myNode
	openList []*myNode
	start    *myNode
	goal     *myNode
	iGraph.Graph
}

type myNode struct {
	tag        string
	neighbors  []*myNode
	next       *myNode
	costToGoal int
	key int
	iGraph.Node
}

func (node myNode) GetValue() int {
	return 1
}

func (pizi *piziStar) Start() {
	// Initial
	processResult := 0
	for processResult > -1 {
		processResult = pizi.process_state()
	}
	// reconstruction()
}

func (pizi *piziStar) process_state() int {
	x := pizi.min_state()
	if x == nil {
		return -1
	}

	k_old := pizi.get_kmin()
	pizi.delete(x)

	if k_old < x.h() {
		for _, neighbor := range x.neighbors {
			if neighbor.h() <= k_old && x.h() > neighbor.h()+pizi.c(neighbor, x) {
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
	return pizi.openList[0]
}

// returns Kmin for the OPENLIST (-1 if the list is empty)
func (pizi *piziStar) get_kmin() int {

	return pizi.min_state().k()
}

func (pizi *piziStar) delete(node *myNode) {
	// METHOD 1
	// var idx int
	// for index, curNode := range pizi.openList {
	// 	if curNode == node {
	// 		idx = index
	// 	} 
	// }

	// method 2
	idx := slices.Index[*myNode](pizi.openList, node)

	pizi.openList[idx].tag = "closed"
	pizi.openList = append(pizi.openList[:idx], pizi.openList[idx+1:]...)
}

func (node *myNode) h() int {
	return node.costToGoal
}

func (node *myNode) b() *myNode {
	return node.next
}

func (pizi *piziStar) c(a, b *myNode) int {
	for _, edge := range pizi.Graph.Edges {
		if edge.Origin == b && edge.Destination == a {
			return edge.Weight
		}
	}
	return -1
}

func (node *myNode) k() int {
	return 1
}

func (pizi *piziStar) Insert(node *myNode, h_new int) {
	kx := -1
	if node.tag == "new" {
		node.tag = "open"
		kx = h_new
		node.key = kx
	} else if node.tag == "open" {
		kx = int(math.Min(float64(node.key), float64(h_new)))
	} else if node.tag == "closed" {
		kx = int(math.Min(float64(node.h()), float64(h_new)))
		node.key = kx
		node.costToGoal = h_new
		node.tag = "open"
	}

	node.costToGoal = h_new
	if len(pizi.openList) > 0 {
		checksum := 0
		for i, curNode := range pizi.openList {
			if node.key <= curNode.key {
				// slices.Insert[*myNode](pizi.openList, i, node)
				pizi.openList = append(pizi.openList[:i+1], pizi.openList[i:]...)
				pizi.openList[i] = node
				break
			} else {
				checksum++
			}
		}
		if checksum == len(pizi.openList) {
			pizi.openList = append(pizi.openList, node)
		}
	} else {
			pizi.openList = append(pizi.openList, node)
	}
}
