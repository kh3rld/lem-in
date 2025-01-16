package internal

import (
	"reflect"
	"testing"
)

func TestBfs(t *testing.T) {
	tests := []struct {
		name         string
		graph        map[string]map[string]int
		startRoom    string
		endRoom      string
		expectedPath []string
	}{
		{"Simple Path", map[string]map[string]int{"start": {"A": 1}, "A": {"end": 1}, "end": {}}, "start", "end", []string{"start", "A", "end"}},
		{"No Path Available", map[string]map[string]int{"start": {"A": 0}, "A": {"end": 0}, "end": {}}, "start", "end", []string{}},
		{"Multiple Possible Paths", map[string]map[string]int{"start": {"A": 1, "B": 1}, "A": {"end": 1}, "end": {}}, "start", "end", []string{"start", "A", "end"}}, // BFS will find this path first
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AntFarm{
				rooms:     make(map[string]*Room),
				startRoom: &Room{name: tt.startRoom},
				endRoom:   &Room{name: tt.endRoom},
			}

			path := af.bfs(tt.graph)
			if !reflect.DeepEqual(path, tt.expectedPath) {
				t.Errorf("bfs() = %v, want %v", path, tt.expectedPath)
			}
		})
	}
}

func TestEdmondsKarp(t *testing.T) {
	tests := []struct {
		name          string
		rooms         map[string]*Room
		startRoom     string
		endRoom       string
		connections   map[string][]string
		expectedPaths [][]string
	}{
		{"Simple Linear Path", map[string]*Room{"start" :{},"A": {},"end": {},}, "start","end", map[string][]string{"start": {"A"}, "A": {"end"},}, [][]string{{"start", "A", "end"},},},
		{"Two Parallel Paths", map[string]*Room{"start": {}, "A" : {}, "B": {}, "end": {},}, "start", "end", map[string][]string{"start": {"A", "B"}, "A": {"end"}, "B": {"end"}, },[][]string{{"start", "A", "end"}, {"start", "B", "end"},},},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build the and farm
			af := &AntFarm{
				rooms: tt.rooms,
				startRoom: tt.rooms[tt.startRoom],
				endRoom: tt.rooms[tt.endRoom],
			}

			// Setuo connections
			for source, destinations ;= range tt.connections {
				sourceRoom := af.rooms[source]
				for _, dest := range desitnations {
					destRoom := af.rooms[dest]
					sourceRoom.connections = append(sourceRoom.connections, destRoom)
				}
			}

			// Run Edmonds-Karp
			af.EdmondsKarp()
		})
	}
}
