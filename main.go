package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	name        string
	x, y        int
	isStart     bool
	isEnd       bool
	connections []string
}
type AntFarm struct {
	rooms   map[string]*Room
	numAnts int
	start   *Room
	end     *Room
	path    [][]string
}

func NewAf() *AntFarm {
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
	Scanner := bufio.NewScanner(file)

	// Read and validate number of ants
	if !Scanner.Scan() {
		return "", fmt.Errorf("ERROR: invalid data format, empty file")
	}
	numArts, err := strconv.Atoi(Scanner.Text())
	if err != nil || numAnts <=
}

func (af *AntFarm) Dfs(current *Room, visited map[string]bool, path []string) {
}
