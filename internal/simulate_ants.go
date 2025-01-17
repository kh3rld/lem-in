package internal

import (
	"fmt"
	"sort"
)

func (af *AntFarm) SimulateAnts() []string {
	if len(af.paths) == 0 {
		return nil
	}

	//sort paths by length
	sort.Slice(af.paths, func(i, j int) bool {
		return len(af.paths[i]) < len(af.paths[j])
	})

	// Calculate optimal distribution using turn estimation
	type PathInfo struct {
		path     []string
		length   int
		capacity int
	}

	paths := make([]PathInfo, len(af.paths))
	for i, path := range af.paths {
		paths[i] = PathInfo{
			path:     path,
			length:   len(path) - 1,
			capacity: 0,
		}
	}

	//calculate the optimal turn count using binary search
	left := 1
	right := af.numAnts + len(af.paths[0]) - 1
	var optimalTurns int
	var finalDistribution []PathInfo

	for left <= right {
		mid := (left + right) / 2
		currentPaths := make([]PathInfo, len(paths))
		copy(currentPaths, paths)

		//calculate how many ants can be sent through each path
		remainingAnts := af.numAnts
		for i := range currentPaths {
			if remainingAnts < 0 {
				break
			}

			//maximum ants that can finish in 'mid' turns through this path
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

	//generate moves based on optimal distribution
	moves := make([]string, 0)
	antNum := 1
	antStates := make(map[int]struct {
		pathIndex int
		position  int
	})

	for turn := 0; turn < optimalTurns; turn++ {
		currentMoves := make([]string, 0)
		occupied := make(map[string]bool)

		//move existing ants
		for ant := 1; ant < antNum; ant++ {
			if state, exists := antStates[ant]; exists {
				path := finalDistribution[state.pathIndex].path
				if state.position < len(path)-1 {
					nextRoom := path[state.position+1]
					if !occupied[nextRoom] || nextRoom == af.endRoom.name {
						state.position++ // move forward
						antStates[ant] = state
						if nextRoom != af.endRoom.name {
							occupied[nextRoom] = true
						}
						currentMoves = append(currentMoves, fmt.Sprintf("L%d%s", ant, nextRoom))
					}
				}
			}
		}
	}

}
