package internal


// Edmonds-Karp algroithm to find augmenting paths
func (af *AntFarm) EdmondsKarp() {
	// Create residual graph
	residualGraph := make(map[string]map[string]int)

	// Inialize residual graph
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
			residualGraph[u][v]--		 // Decrease forward edge
			if residualGraph[v][u] == 0 { // Add reverse edge if it doesn't exist
				residualGraph[v][u] = 1
			}
		}
	}
}

// bfs implements breath-first search to find shortest augmenting path
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
					// Construct Path
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