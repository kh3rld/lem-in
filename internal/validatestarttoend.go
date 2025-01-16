package internal

import "fmt"

func NewPathValidation() *PathValidation {
	return &PathValidation{
		visited: make(map[string]bool),
	}
}
func (af *AntFarm) ValidateStartEndPath() error {
	// Initialize path validation
	pv := NewPathValidation()

	// Check if path exists using DFS
	if !pv.hasPath(af.startRoom, af.endRoom) {
		return fmt.Errorf("ERROR: invalid data format, no path exists between start and end rooms")
	}
	return nil
}

// hasPath performs DFS to check if there's a path between start and end rooms
func (pv *PathValidation) hasPath(current *Room, end *Room) bool {
	// If we've reached the end room, we found a path
	if current == end {
		return true
	}

	// Mark current room as visited
	pv.visited[current.name] = true

	// Check all connections from current room
	for _, nextRoom := range current.connections {
		// If room hasn't been visited, explore it
		if !pv.visited[nextRoom.name] {
			if pv.hasPath(nextRoom, end) {
				return true
			}
		}
	}
	return false
}
