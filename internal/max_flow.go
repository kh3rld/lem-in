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

	
}

// Bfs implements breath-first search to find shortest augmenting path
func (af *AntFarm) bfs(residualGraph map[string]map[string]int) []string {
	// function signature
	return []string{}
}