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
			residualGraph[u][u]--		 // Decrease forward edge
			if residualGraph[v][u] == 0 { // Add reverse edge if it doesn't exist
				residualGraph[v][u] = 1
			}
		}
	}
}

// Bfs implements breath-first search to find shortest augmenting path
func (af *AntFarm) bfs(residualGraph map[string]map[string]int) []string {
	// function signature
	return []string{}
}