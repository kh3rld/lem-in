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

	//parse rooms and links
	parsingRooms := true
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		//handling special commands
		if line == "##start" || line == "##end" {
			if !scanner.Scan() {
				return fmt.Errorf("ERROR: invalid data format, missing room after %s", line)
			}
			roomLine := scanner.Text()
			if err := af.ParseRoom(roomLine, line == "##start"); err != nil {
				return err
			}
			continue
		}

		//If line contains a -, parse the links
		if strings.Contains(line, "-") {
			if parsingRooms {
				parsingRooms = false
			}
			if err := af.Parselink(line); err != nil {
				return err
			}
		}
	}
	return nil
}

func (af *AntFarm) ParseRoom (line string, isStart bool) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("ERROR: invalid data format, invalid room name")
	}

	name := parts[0]
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return fmt.Errorf("ERROR: invalid data format, invalid room name")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf(("ERROR: invalid data format, invalid x co-ordinate"))
	}

	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format, invalid y co-ordinate")
	}

	room :=  &Room {
		name: name,
		x: x,
		y: y,
		isStart: isStart,
		connections: make([]*Room, 0),
	}
	if isStart {
		if af.startRoom != nil {
			return fmt.Errorf("ERROR: invalid data format, multipla start rooms")
		}
		af.startRoom = room
	}
	af.rooms[name] = room
	return nil
}

//parslink parses a link line and adds the conections
func (af *AntFarm) Parselink(line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("ERROR: invalid data format, invalid link format")
	}

	room1, exists1 := af.rooms[parts[0]]
	room2, exists2 := af.rooms[parts[1]]

	if !exists1 || !exists2 {
		return  fmt.Errorf("ERROR: invalid data format, link to unknown room")
	}
}