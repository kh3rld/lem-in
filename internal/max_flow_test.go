package internal

import (
	"testing"
)

func TestEdmondsKarp(t *testing.T) {
	tests := []struct {
		name     string
		input    *AntFarm
		wantLen  int  // Expected number of paths
		wantPath bool // Whether specific paths should be checked
	}{
		{
			name: "Simple path",
			input: &AntFarm{
				rooms: map[string]*Room{
					"start": {name: "start", connections: []*Room{}},
					"1":     {name: "1", connections: []*Room{}},
					"end":   {name: "end", connections: []*Room{}},
				},
				startRoom: &Room{name: "start"},
				endRoom:   &Room{name: "end"},
			},
			wantLen:  1,
			wantPath: true,
		},
		{
			name: "No path",
			input: &AntFarm{
				rooms: map[string]*Room{
					"start": {name: "start", connections: []*Room{}},
					"end":   {name: "end", connections: []*Room{}},
				},
				startRoom: &Room{name: "start"},
				endRoom:   &Room{name: "end"},
			},
			wantLen:  0,
			wantPath: false,
		},
		{
			name: "Multiple paths",
			input: &AntFarm{
				rooms: map[string]*Room{
					"start": {name: "start", connections: []*Room{}},
					"a":     {name: "a", connections: []*Room{}},
					"b":     {name: "b", connections: []*Room{}},
					"c":     {name: "c", connections: []*Room{}},
					"d":     {name: "d", connections: []*Room{}},
					"end":   {name: "end", connections: []*Room{}},
				},
				startRoom: &Room{name: "start"},
				endRoom:   &Room{name: "end"},
			},
			wantLen:  2,
			wantPath: true,
		},
	}

	// Set up connections for test cases
	tests[0].input.rooms["start"].connections = []*Room{tests[0].input.rooms["1"]}
	tests[0].input.rooms["1"].connections = []*Room{tests[0].input.rooms["end"]}

	// Create a graph with two possible disjoint paths:
	// start -> a -> c -> end
	// start -> b -> d -> end
	tests[2].input.rooms["start"].connections = []*Room{
		tests[2].input.rooms["a"],
		tests[2].input.rooms["b"],
	}
	tests[2].input.rooms["a"].connections = []*Room{tests[2].input.rooms["c"]}
	tests[2].input.rooms["b"].connections = []*Room{tests[2].input.rooms["d"]}
	tests[2].input.rooms["c"].connections = []*Room{tests[2].input.rooms["end"]}
	tests[2].input.rooms["d"].connections = []*Room{tests[2].input.rooms["end"]}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.input.EdmondsKarp()

			// Check number of paths found
			if got := len(tt.input.paths); got != tt.wantLen {
				t.Errorf("EdmondsKarp() path count = %v, want %v", got, tt.wantLen)
			}

			if tt.wantPath {
				// Verify path properties
				for _, path := range tt.input.paths {
					// Check if path starts and ends correctly
					if len(path) > 0 {
						if path[0] != tt.input.startRoom.name {
							t.Errorf("Path doesn't start at start room: %v", path)
						}
						if path[len(path)-1] != tt.input.endRoom.name {
							t.Errorf("Path doesn't end at end room: %v", path)
						}
					}

				
	}
}
