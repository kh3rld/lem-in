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

	
}
