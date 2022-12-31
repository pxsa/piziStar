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
	// شب که میشه به عشق تو غزل غزل صدا میشم ترانه خون قضه تمام عاشقا میشم
	l.nodes = make([][]*piziNode, rows)
	for i := range l.nodes {
		l.nodes[i] = make([]*piziNode, cols)

		for j := range l.nodes[i] {
			l.nodes[i][j] = initNode(i, j)
		}
	}
}

func initNode(i, j int) *piziNode{
	return &piziNode{
		row: i,
		col: j,
		tag: "new",
	}
}

func (l *Land) detectNeighbors(node *piziNode) {
	neighbors := []*piziNode{}

	// upper neighbor
	if node.row - 1 >= 0 {
		neighbors = append(neighbors, l.nodes[node.row-1][node.col])
	}
	
	// bottom neighbor
	if node.row + 1 < l.rows {
		neighbors = append(neighbors, l.nodes[node.row+1][node.col])
	}
	
	// left neighbor
	if node.col - 1 >= 0 {
		neighbors = append(neighbors, l.nodes[node.row][node.col-1])
	}
	
	// right neighbor
	if node.col + 1 < l.cols {
		neighbors = append(neighbors, l.nodes[node.row][node.col+1])
	}

	node.neighbors = neighbors
}


func (l *Land) print() {
	fmt.Println(l.nodes)
}