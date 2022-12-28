package main

import (
	"github.com/pxsa/iGraph"
)

func main() {

	goal := &myNode{
		tag: "new",
		Node: iGraph.Node{
			Value: 2,
		},
		costToGoal: 0,
		next: nil,
	}
	start := &myNode{
		tag: "new",
		
	}

	state1 := &myNode{
		tag: "new",
	}
	state2 := &myNode{
		tag: "new",
	}
	state3 := &myNode{
		tag: "new",
	}
	states := []iGraph.CustomNode {state1, state2, state3}

	edge1 := &iGraph.Edge {
		Weight: 5,
		Origin: start,
		Destination: state1,
		IsDirected: true,
	}
	edge2 := &iGraph.Edge {
		Weight: 5,
		Origin: state1,
		Destination: goal,
		IsDirected: true,
	}
	
	openList := []*myNode{}
	openList = append(openList, goal)
	
	edges := []*iGraph.Edge {edge1, edge2}
	// graph := iGraph.Graph {
	// 	Nodes: states,
	// 	Edges: edges,
	// }

	pizi := piziStar{
		start: start,
		goal:  goal,
		openList: openList,
		states: []*myNode{state1, state2, state3},
	}

	pizi.process_state()

}
