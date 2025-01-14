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

func (af *AntFarm) ParseInput(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var fileContent strings.Builder
	scanner := bufio.NewScanner(file)

	// Read and validate number of ants
	if !scanner.Scan() {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	numAnts, err := strconv.Atoi(scanner.Text())
	if err != nil || numAnts <= 0 {
		return "", fmt.Errorf("ERROR: invalid data format, empty file")
	}
	af.numAnts = numAnts
	fileContent.WriteString(fmt.Sprintf("%d\n", numAnts))

	parsingRooms := true
	var isStart, isEnd bool

	for scanner.Scan() {
		line := scanner.Text()
		fileContent.WriteString(line + "\n")

		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}
		
		if line == "##end" {
			isEnd = true
			continue
		}

		if strings.Contains(line, "-") {
			if parsingRooms
		}
	}
}
