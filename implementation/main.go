package main

func main() {

	goal := myNode{
		tag: "new",
	}
	start := myNode{
		tag: "new",
	}

	state1 := myNode{
		tag: "new",
	}
	state2 := myNode{
		tag: "new",
	}
	state3 := myNode{
		tag: "new",
	}
	states := []myNode {state1, state2, state3}

	openList := []myNode{}
	openList = append(openList, goal)

	pizi := piziStar{
		start: start,
		goal:  goal,
		openList: openList,
		states: states,
	}
	
	pizi.process_state()

}
