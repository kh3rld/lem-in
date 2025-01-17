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
}
