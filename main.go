package main

import (
	"fmt"
	"os"
)

// Room represent a node in the ant farm
type Room struct {
	name        string
	x, y        int
	isStart     bool
	isEnd       bool
	connections []*Room
}

// Ant farm represents the whole colony
type AntFarm struct {
	rooms     map[string]*Room
	startRoom *Room
	endRoom   *Room
	numAnts   int
	paths     [][]string
}

// initialise a new ant farm
func NewAntFarm() *AntFarm {
	return &AntFarm{
		rooms: make(map[string]*Room),
	}
}

// validate the input file
func (af *AntFarm) ParseInput(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	return nil
}
