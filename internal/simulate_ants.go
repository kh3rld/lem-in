package internal

import (
	"fmt"
	"sort"
	"strings"
)

// PathInfo holds information about each path, including the path itself,
// its length, and the capacity of ants that can use it.
type PathInfo struct {
	path     []string // List of rooms in the path
	length   int      // Number of rooms in the path (excluding start and end)
	capacity int      // Number of ants that can currently be assigned to this path
}

func (af *AntFarm) SimulateAnts() []string {
	if len(af.paths) == 0 {
		return nil
	}

	// Sort paths by length
	sort.Slice(af.paths, func(i, j int) bool {
		return len(af.paths[i]) < len(af.paths[j])
	})

	// Calculate optimal distribution of ants
	paths := calculatePathsInfo(af.paths)
	optimalTurns, finalDistribution := findOptimalTurns(paths, af.numAnts)

	// Generate and return moves
	return generateMoves(finalDistribution, optimalTurns, af.numAnts, af.endRoom.name)
}


func calculatePathsInfo(paths [][]string) []PathInfo {
	pathsInfo := make([]PathInfo, len(paths))
	for i, path := range paths {
		pathsInfo[i] = PathInfo{
			path:     path,
			length:   len(path) - 1,
			capacity: 0,
		}
	}
	return pathsInfo
}

func findOptimalTurns(paths []PathInfo, numAnts int) (int, []PathInfo) {
	left, right := 1, numAnts+len(paths[0].path)-1
	var optimalTurns int
	var finalDistribution []PathInfo

	for left <= right {
		mid := (left + right) / 2
		currentPaths := make([]PathInfo, len(paths))
		copy(currentPaths, paths)

		// Calculate ant distribution for the current turn count (mid)
		remainingAnts := numAnts
		for i := range currentPaths {
			if remainingAnts < 0 {
				break
			}

			// Calculate the maximum number of ants that can be sent through this path in 'mid' turns
			maxAnts := mid - currentPaths[i].length + 1
			if maxAnts > 0 {
				antsToSend := min(remainingAnts, maxAnts)
				currentPaths[i].capacity = antsToSend
				remainingAnts -= antsToSend
			}
		}

		if remainingAnts <= 0 {
			optimalTurns = mid
			finalDistribution = make([]PathInfo, len(currentPaths))
			copy(finalDistribution, currentPaths)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return optimalTurns, finalDistribution
}

func generateMoves(paths []PathInfo, optimalTurns, numAnts int, endRoomName string) []string {
	moves := make([]string, 0)
	antNum := 1
	antStates := make(map[int]struct {
		pathIndex int
		position  int
	})

	for turn := 0; turn < optimalTurns; turn++ {
		currentMoves := make([]string, 0)
		occupied := make(map[string]bool)

		// Move existing ants
		moveExistingAnts(&antStates, paths, occupied, &currentMoves, endRoomName)

		// Start new ants
		startNewAnts(paths, &antStates, &antNum, occupied, &currentMoves)

		if len(currentMoves) > 0 {
			sort.Strings(currentMoves)
			moves = append(moves, strings.Join(currentMoves, " "))
		}
	}

	return moves
}

func moveExistingAnts(antStates *map[int]struct {
	pathIndex int
	position  int
}, paths []PathInfo, occupied map[string]bool, currentMoves *[]string, endRoomName string) {
	for ant, state := range *antStates {
		path := paths[state.pathIndex].path
		if state.position < len(path)-1 {
			nextRoom := path[state.position+1]
			if !occupied[nextRoom] || nextRoom == endRoomName {
				// Move ant forward
				state.position++
				(*antStates)[ant] = state
				if nextRoom != endRoomName {
					occupied[nextRoom] = true
				}
				*currentMoves = append(*currentMoves, fmt.Sprintf("L%d-%s", ant, nextRoom))
			}
		}
	}
}

func startNewAnts(paths []PathInfo, antStates *map[int]struct {
	pathIndex int
	position  int
}, antNum *int, occupied map[string]bool, currentMoves *[]string) {
	for i := range paths {
		if paths[i].capacity > 0 {
			nextRoom := paths[i].path[1]
			if !occupied[nextRoom] {
				(*antStates)[*antNum] = struct {
					pathIndex int
					position  int
				}{i, 1}
				occupied[nextRoom] = true
				*currentMoves = append(*currentMoves, fmt.Sprintf("L%d-%s", *antNum, nextRoom))
				*antNum++ // Increment ant number
				paths[i].capacity--
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
