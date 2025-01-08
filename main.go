package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return fmt.Errorf("ERROR: invalid data format, empty file")
	}
	numAnts, err := strconv.Atoi(scanner.Text())
	if err != nil || numAnts == 0 {
		return fmt.Errorf("ERROR: invalid data format, invalid number of ants")
	}
	af.numAnts = numAnts

	return nil
}
