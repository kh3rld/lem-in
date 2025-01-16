package edmondskarp

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
