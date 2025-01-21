package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		return "", fmt.Errorf("ERROR: invalid data format, empty file")
	}
	numAnts, err := strconv.Atoi(scanner.Text())
	if err != nil || numAnts <= 0 {
		return "", fmt.Errorf("ERROR: invalid data format, invalid number of ants")
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

		if line == "##start" {
			isStart = true
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
	Reading(fileContent.String())
	if af.startRoom == nil {
		return "", fmt.Errorf("ERROR: invalid data format, no start room found")
	}
	if af.endRoom == nil {
		return "", fmt.Errorf("ERROR: invalid data format, no end room found")
	}
	
	startToEnd := af.ValidateStartEndPath()
	if startToEnd != nil {
		return "", fmt.Errorf("ERROR: invalid data format, no link from start to end")
	}

	return fileContent.String(), nil
}

