package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

// AntFarm represents the whole colony
type AntFarm struct {
	rooms     map[string]*Room
	startRoom *Room
	endRoom   *Room
	numAnts   int
	paths     [][]string
}

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

	if af.startRoom == nil {
		return "", fmt.Errorf("ERROR: invalid data format, no start room found")
	}
	if af.endRoom == nil {
		return "", fmt.Errorf("ERROR: invalid data format, no end room found")
	}

	return fileContent.String(), nil
}

func (af *AntFarm) ParseRoom(line string, isStart bool, isEnd bool) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("ERROR: invalid data format, invalid room definition")
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

func (af *AntFarm) Parselink(line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("ERROR: invalid data format, invalid link format")
	}

	room1, exists1 := af.rooms[parts[0]]
	room2, exists2 := af.rooms[parts[1]]

	if !exists1 || !exists2 {
		return fmt.Errorf("ERROR: invalid data format, link to unknown room")
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

// EdmondsKarp implements the Edmonds-Karp algorithm to find augmenting paths
func (af *AntFarm) EdmondsKarp() {
	// Create residual graph
	residualGraph := make(map[string]map[string]int)

	// Initialize residual graph
	for name, room := range af.rooms {
		residualGraph[name] = make(map[string]int)
		for _, conn := range room.connections {
			residualGraph[name][conn.name] = 1 // Initial capacity of 1 for each edge
		}
	}

	af.paths = make([][]string, 0)

	// Keep finding paths until no more paths exist
	for {
		path := af.bfs(residualGraph)
		if len(path) == 0 {
			break
		}
		af.paths = append(af.paths, path)

		// Update residual graph
		for i := 0; i < len(path)-1; i++ {
			u, v := path[i], path[i+1]
			residualGraph[u][v]--         // Decrease forward edge
			if residualGraph[v][u] == 0 { // Add reverse edge if it doesn't exist
				residualGraph[v][u] = 1
			}
		}
	}
}

// bfs implements breadth-first search to find shortest augmenting path
func (af *AntFarm) bfs(residualGraph map[string]map[string]int) []string {
	visited := make(map[string]bool)
	parent := make(map[string]string)
	queue := []string{af.startRoom.name}
	visited[af.startRoom.name] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for next, capacity := range residualGraph[current] {
			if !visited[next] && capacity > 0 {
				visited[next] = true
				parent[next] = current
				queue = append(queue, next)

				if next == af.endRoom.name {
					// Construct path
					path := []string{next}
					for p := current; p != af.startRoom.name; p = parent[p] {
						path = append([]string{p}, path...)
					}
					path = append([]string{af.startRoom.name}, path...)
					return path
				}
			}
		}
	}
	return []string{}
}

func (af *AntFarm) Dfs(current *Room, visited map[string]bool, path []string) {
	visited[current.name] = true
	path = append(path, current.name)

	if current == af.endRoom {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		af.paths = append(af.paths, pathCopy)
	} else {
		for _, next := range current.connections {
			if !visited[next.name] {
				af.Dfs(next, visited, path)
			}
		}
	}
	visited[current.name] = false
}

func (af *AntFarm) SimulateAnts() []string {
    if len(af.paths) == 0 {
        return nil
    }

    // Sort paths by length
    sort.Slice(af.paths, func(i, j int) bool {
        return len(af.paths[i]) < len(af.paths[j])
    })

    // Calculate flow distribution
    type PathFlow struct {
        path     []string
        antCount int
        length   int
    }

    // Initialize path flows
    flows := make([]PathFlow, len(af.paths))
    for i, path := range af.paths {
        flows[i] = PathFlow{
            path:     path,
            antCount: 0,
            length:   len(path) - 1, // -1 because we don't count start room
        }
    }

    // Distribute ants optimally
    remainingAnts := af.numAnts
    for remainingAnts > 0 {
        // Find best path for next ant
        bestPath := -1
        minTotalTime := math.MaxInt32

        for i := range flows {
            // Calculate completion time if we add an ant to this path
            completionTime := flows[i].length + flows[i].antCount
            if completionTime < minTotalTime {
                minTotalTime = completionTime
                bestPath = i
            }
        }

        if bestPath == -1 {
            break
        }

        flows[bestPath].antCount++
        remainingAnts--
    }

    // Simulate movements
    moves := make([]string, 0)
    type AntState struct {
        pathIndex  int
        position   int
        completed  bool
    }
    antStates := make(map[int]*AntState)
    currentAnt := 1

    // Calculate maximum possible turns
    maxTurns := flows[0].length + af.numAnts

    for turn := 0; turn < maxTurns; turn++ {
        currentTurn := make([]string, 0)
        roomOccupied := make(map[string]bool)

        // Move existing ants first
        for ant := 1; ant <= currentAnt-1; ant++ {
            state, exists := antStates[ant]
            if !exists || state.completed {
                continue
            }

            path := flows[state.pathIndex].path
            if state.position < len(path)-1 {
                nextRoom := path[state.position+1]
                if !roomOccupied[nextRoom] || nextRoom == af.endRoom.name {
                    state.position++
                    if nextRoom != af.endRoom.name {
                        roomOccupied[nextRoom] = true
                    }
                    currentTurn = append(currentTurn,
                        fmt.Sprintf("L%d-%s", ant, nextRoom))
                    if state.position == len(path)-1 {
                        state.completed = true
                    }
                }
            }
        }

        // Try to start new ants
        for i := range flows {
            if flows[i].antCount > 0 {
                nextRoom := flows[i].path[1] // First room after start
                if !roomOccupied[nextRoom] {
                    antStates[currentAnt] = &AntState{
                        pathIndex: i,
                        position:  1,
                        completed: false,
                    }
                    roomOccupied[nextRoom] = true
                    currentTurn = append(currentTurn,
                        fmt.Sprintf("L%d-%s", currentAnt, nextRoom))
                    currentAnt++
                    flows[i].antCount--
                }
            }
        }

        if len(currentTurn) > 0 {
            sort.Strings(currentTurn)
            moves = append(moves, strings.Join(currentTurn, " "))
        }

        // Check if all ants have reached the end
        allDone := true
        for ant := 1; ant <= af.numAnts; ant++ {
            state, exists := antStates[ant]
            if !exists || !state.completed {
                allDone = false
                break
            }
        }
        if allDone {
            break
        }
    }

    return moves
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [filename]")
		return
	}

	farm := NewAntFarm()
	content, err := farm.ParseInput(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the input file content
	fmt.Print(content)

	// Print a blank line before ant movements
	fmt.Println()

	// Find paths and simulate ant movements
	farm.EdmondsKarp()
	moves := farm.SimulateAnts()

	// Print ant movements
	for _, move := range moves {
		fmt.Println(move)
	}
}
