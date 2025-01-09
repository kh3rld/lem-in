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

func (af *AntFarm) FindPaths() {
    visited := make(map[string]bool)
    currentPath := make([]string, 0)
    af.Dfs(af.startRoom, visited, currentPath)
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

    // Find shortest path
    shortestPath := af.paths[0]
    for _, path := range af.paths {
        if len(path) < len(shortestPath) {
            shortestPath = path
        }
    }

    moves := make([]string, 0)
    antPositions := make(map[int]int) // ant number -> position in path
    antNum := 1

    for len(antPositions) > 0 || antNum <= af.numAnts {
        currentTurn := make([]string, 0)

        // Move existing ants
        for ant := 1; ant < antNum; ant++ {
            if pos, exists := antPositions[ant]; exists {
                if pos < len(shortestPath)-1 {
                    antPositions[ant]++
                    currentTurn = append(currentTurn, fmt.Sprintf("L%d-%s",
                        ant, shortestPath[antPositions[ant]]))
                } else {
                    delete(antPositions, ant)
                }
            }
        }

        // Add new ant if possible
        if antNum <= af.numAnts {
            antPositions[antNum] = 0
            currentTurn = append(currentTurn, fmt.Sprintf("L%d-%s", antNum, shortestPath[0]))
            antNum++
        }

        if len(currentTurn) > 0 {
            moves = append(moves, strings.Join(currentTurn, " "))
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
    farm.FindPaths()
    moves := farm.SimulateAnts()
    
    // Print ant movements
    for _, move := range moves {
        fmt.Println(move)
    }
}