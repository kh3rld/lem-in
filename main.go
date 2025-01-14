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
			if parsingRooms {
				parsingRooms = false
			}
			if err := af.Parselink(line); err != nil {
				return "", err
			}
			continue
		}

		if parsingRooms && len(line) > 0 {
			if err := af.ParseRoom(line, isStart, isEnd); err != nil {
				return "", err
			}
			isStart = false
			isEnd = false
		}
	}

	if af.startRoom == nil {
		return "", fmt.Errorf("ERROR: invalid data format, no end room found")
	}
	if af.endRoom == nil {
		return "", fmt.Errorf("ERROR: invalid data format, no end room found")
	}

	return fileContent.String(), nil
}

// Validates and Parses Room data into the relevant data structures
func (af *AntFarm) ParseRoom(line string, isStart bool, isEnd bool) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("ERROR: invalid data format, invalid room defination")
	}

	name := parts[0]
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return fmt.Errorf("ERROR: invalid data format, invalid room name")
	}

	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format, invalid x coordinate")
	}

	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format, invalid y coordinate")
	}

	room := &Room{
		name:        name,
		x:           x,
		y:           y,
		isStart:     isStart,
		isEnd:       isEnd,
		connections: make([]*Room, 0),
	}
	if isStart {
		if af.startRoom != nil {
			return fmt.Errorf("ERROR: invalid data format, multiple start rooms")
		}
		af.startRoom = room
	}
	
	if isEnd {
		if af.endRoom != nil {
			return fmt.Errorf("ERROR: invalid data format, multiple end rooms")
		}
		af.endRoom = room
	}

	af.rooms[name] = room
	return nil
}

// Validates and Parses data of how rooms are linked
func (af *AntFarm) Parselink(line string) error {
	parts := strings.Split(line, "-") 
	if len(parts) != 2 {
		return fmt.Errorf("ERROR: invalid data format, invalid link format")
	}

	room1, exists1 := af.rooms[parts[0]]
	room2, exists2 := af.rooms[parts[1]]

	if !exists1 || !exists2 {
		return fmt.Errorf("ERROR: invalid data format, link to unkown room")
	}

	for _, conn := range room1.connections {
		if conn == room2 {
			return fmt.Errorf("ERROR: invalid data format, duplicate link")
		}
	}

	room1.connections = append(room1.connections, room2)
	room2.connections = append(room2.connections, room1)
	return nil

}

// EdmondsKarp implements the Edmonds-Karp algorithm to find augmnenting paths
func (af *AntFarm) EdmondsKarp() {
	// Implementation
}
