package main

// Room represent a node in the ant farm
type Room struct {
	name        string
	x, y        int
	isStart     bool
	isEnd       bool
	connections []*Room
}

//Ant farm represents the whole colony 
type AntFarm struct {
	rooms map[string]*Room
	startRoom *Room
	endRoom	*Room
	numAnts int
	paths [][]string
}

//initialise a new ant farm
func NewAntFarm() *AntFarm {
	return &AntFarm{
		rooms: make(map[string]*Room),
	}
}