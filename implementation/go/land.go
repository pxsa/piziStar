package main

import (
	"fmt"
)

type Land struct {
	rows int 
	cols int 
	nodes [][] *piziNode
}

func (l *Land) init(rows, cols int) {
	l.rows = rows
	l.cols = cols
	// (ابی)شب که میشه به عشق تو غزل غزل صدا میشم ترانه خون قصه تمام عاشقا میشم
	l.nodes = make([][]*piziNode, rows)
	for i := range l.nodes {
		l.nodes[i] = make([]*piziNode, cols)

		for j := range l.nodes[i] {
			l.nodes[i][j] = initNode(i, j)
		}
	}

	// this is not good (search for the bether solution)
	for i := range l.nodes {
		for j := range l.nodes[i] {
			l.detectNeighbors(l.nodes[i][j])
		}
	}
}

func (l *Land) makeObstacle(node *piziNode ,status bool) {
	node.isObstacle = status
}

// func (l *Land) makeObstacle(row, col int ,status bool) {
// 	l.nodes[row][col].isObstacle = status
// }

func initNode(i, j int) *piziNode{
	return &piziNode{
		row: i,
		col: j,
		tag: "new",
		isObstacle: false,
		next: nil,
		costToGoal: -1, // not sure about this.
	}
}

func (l *Land) detectNeighbors(node *piziNode) {
	neighbors := []*piziNode{}

	// left neighbor
	if node.row - 1 >= 0 {
		neighbors = append(neighbors, l.nodes[node.row-1][node.col])
	}
	
	// right neighbor
	if node.row + 1 < l.rows {
		neighbors = append(neighbors, l.nodes[node.row+1][node.col])
	}
	
	// upper neighbor
	if node.col - 1 >= 0 {
		neighbors = append(neighbors, l.nodes[node.row][node.col-1])
	}
	
	// lower neighbor
	if node.col + 1 < l.cols {
		neighbors = append(neighbors, l.nodes[node.row][node.col+1])
	}

	node.neighbors = neighbors
}


func (l *Land) print() {
	fmt.Println(l.nodes)
}