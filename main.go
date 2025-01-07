package main

// Room represent a node in the ant farm
type Room struct {
	name        string
	x, y        int
	isStart     bool
	isEnd       bool
	connections []*Room
}
