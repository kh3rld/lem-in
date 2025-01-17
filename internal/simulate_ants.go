package internal

import "sort"

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

	paths := make ([]PathInfo, len(af.paths))
	for i, path := range af.paths {
		paths[i] = PathInfo {
			path: path,
			length: len(path)- 1,
			capacity: 0,
		}
	}

	//calculate the optimal turn count using binary search
	left := 1
	right := af.numAnts+len(af.paths[0]) -1
	var optimalTurns int
	var finalDistribution []PathInfo

	for left <= right {
		mid := (left + right) / 2
		currentPaths := make([]PathInfo, len(paths))
		copy(currentPaths, paths)
	}
}
