package main

import (
	"fmt"
	"math"
	"sort"

	"golang.org/x/exp/slices"
)

type piziStar struct {
	start    *piziNode
	goal     *piziNode
	openList []*piziNode
}

func (pizi *piziStar) Start() {
	// Initial
	pizi.Insert(pizi.goal, 0)
	processResult := 0

	for processResult > -1 {
		processResult = pizi.process_state()
	}
	path := pizi.reconstruction()
	showPath(path)
}

func showPath(path []*piziNode) {
	fmt.Println("")
}

func (pizi *piziStar) reconstruction() []*piziNode {
	path := []*piziNode{}
	tmp := pizi.start
	for tmp != pizi.goal {
		path = append(path, tmp)
		tmp = tmp.next
	}
	path = append(path, pizi.goal)
	return path
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
		for _, y := range x.neighbors {
			if (y.tag == "new") ||
			 (y.next == x && y.costToGoal != x.costToGoal+pizi.c(x, y)) || 
			 (y.next != x && y.costToGoal > x.costToGoal +pizi.c(x, y)) {
				y.next = x
				new_h := x.costToGoal + pizi.c(x, y)
				pizi.Insert(y, new_h)
			}
		}

	} else {
		for _, y := range x.neighbors {
			if (y.tag == "new") || (y.next == x && y.costToGoal != x.costToGoal+pizi.c(x, y)) {
				y.next = x
				new_h := x.costToGoal + pizi.c(x, y)
				pizi.Insert(y, new_h)

			} else {
				if y.next != x && (y.costToGoal > x.costToGoal+pizi.c(x, y)) {
					pizi.Insert(x, x.costToGoal)
				} else {
					if y.next != x && (x.costToGoal > y.costToGoal+pizi.c(y, x)) && y.tag == "closed" && y.costToGoal > k_old {
						pizi.Insert(y, y.costToGoal)
					}
				}
			}
		}
	}

	return pizi.get_kmin()
}

func (pizi *piziStar) modify_cost(x *piziNode, newCost int) int {
	// c(x, y) = cval
	if x.tag == "closed" {
		pizi.Insert(x, newCost)
	}
	return pizi.get_kmin()
}

// returns the state on the OPENLIST with minimum k(.) value
// (NULL if the list is empty)
func (pizi *piziStar) min_state() *piziNode {
	if len(pizi.openList) > 0 {
		return pizi.openList[0]
	}
	return nil
}

// returns Kmin for the OPENLIST (-1 if the list is empty)
func (pizi *piziStar) get_kmin() int {

	if len(pizi.openList) > 0 {
		return pizi.openList[0].key
	}
	return -1
}

func (pizi *piziStar) delete(node *piziNode) {

	idx := slices.Index[*piziNode](pizi.openList, node)

	pizi.openList[idx].tag = "closed"
	pizi.openList = append(pizi.openList[:idx], pizi.openList[idx+1:]...)
}

func (pizi *piziStar) c(a, b *piziNode) int {
	// for _, edge := range pizi.Graph.Edges {
	// 	if edge.Origin == b && edge.Destination == a {
	// 		return edge.Weight
	// 	}
	// }
	return 1
}

func (pizi *piziStar) Insert(node *piziNode, h_new int) {
	kx := -1
	// insertedList.add(node)
	if node.tag == "new" {
		kx = h_new

	} else if node.tag == "open" {
		kx = int(math.Min(float64(node.key), float64(h_new)))

	} else if node.tag == "closed" {
		kx = int(math.Min(float64(node.costToGoal), float64(h_new)))
		//obstacle
	}
	node.tag = "open"
	node.key = kx
	node.costToGoal = h_new

	pizi.openList = append(pizi.openList, node)
	pizi.openList = sortOpenlist(pizi.openList)
}

func sortOpenlist(list []*piziNode) []*piziNode {
	sort.Slice(list, func(i, j int) bool {
		return list[i].key < list[j].key
	})

	return list
}
