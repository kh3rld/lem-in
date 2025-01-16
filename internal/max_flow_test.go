package internal

import (
	"reflect"
	"testing"
)

func Testbfs(t *testing.T) {
}

func TestEdmondsKarp(t *testing.T) {
	tests := []struct {
		name string
		graph map[string]map[string]int
		startRoom string
		endRoom	string
		expectedPath []string
	}{
		{"Simple Path", map[string]map[string]int{"start": {"A": 1}, "A": {"end": 1}, "end": {},}, "start", "end", []string{"start", "A", "end"},},
		{"No Path Available", map[string]map[string]int{"start": {"A": 0}, "A": {"end": 0}, "end": {},}, "start", "end", []string{},},
		{"Multiple Possible Paths", map[string]map[string]int{"start": {"A": 1, "B": 1}, "A": {"end": 1}, "end": {},}, "start", "end", []string{"start", "A", "end"}, }, // BFS will find this path first
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AntFarm{
				rooms: make(map[string]*Room),
				startRoom: &Room{name: tt.startRoom},
				endRoom: &Room{name: tt.endRoom},
			}

			path := af.bfs(tt.graph)
			if !reflect.DeepEqual(path, tt.expectedPath) {
				t.Errorf("bfs() = %v, want %v", path, tt.expectedPath)
			}
		})
	}
	
}
