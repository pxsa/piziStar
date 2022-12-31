package main

func main() {

	// Creating land 
	land := Land{}
	land.init(2, 3)

	

	pizi := piziStar{
		start: land.nodes[0][0],
		goal:  land.nodes[1][2],
	}

	pizi.Start()
	
}