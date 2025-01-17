package internal

import (
	"reflect"
	"testing"
)

func TestCalculatePathsInfo(t *testing.T) {
	tests := []struct {
		name     string
		paths    [][]string
		expected []PathInfo
	}{
		{
			name: "Single path",
			paths: [][]string{
				{"start", "room1", "end"},
			},
			expected: []PathInfo{
				{
					path:     []string{"start", "room1", "end"},
					length:   2,
					capacity: 0,
				},
			},
		},
		{
			name: "Multiple paths",
			paths: [][]string{
				{"start", "room1", "end"},
				{"start", "room2", "room3", "end"},
			},
			expected: []PathInfo{
				{
					path:     []string{"start", "room1", "end"},
					length:   2,
					capacity: 0,
				},
				{
					path:     []string{"start", "room2", "room3", "end"},
					length:   3,
					capacity: 0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculatePathsInfo(tt.paths)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("calculatePathsInfo() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindOptimalTurns(t *testing.T) {
	tests := []struct {
		name            string
		paths           []PathInfo
		numAnts         int
		expectedTurns   int
		expectedDistrib []PathInfo
	}{
		{
			name: "Single path, one ant",
			paths: []PathInfo{
				{
					path:     []string{"start", "room1", "end"},
					length:   2,
					capacity: 0,
				},
			},
			numAnts:       1,
			expectedTurns: 2,
			expectedDistrib: []PathInfo{
				{
					path:     []string{"start", "room1", "end"},
					length:   2,
					capacity: 1,
				},
			},
		},
		{
			name: "Two paths, three ants",
			paths: []PathInfo{
				{
					path:     []string{"start", "room1", "end"},
					length:   2,
					capacity: 0,
				},
				{
					path:     []string{"start", "room2", "end"},
					length:   2,
					capacity: 0,
				},
			},
			numAnts:       3,
			expectedTurns: 3,
			expectedDistrib: []PathInfo{
				{
					path:     []string{"start", "room1", "end"},
					length:   2,
					capacity: 2,
				},
				{
					path:     []string{"start", "room2", "end"},
					length:   2,
					capacity: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			turns, distrib := findOptimalTurns(tt.paths, tt.numAnts)
			if turns != tt.expectedTurns {
				t.Errorf("findOptimalTurns() turns = %v, want %v", turns, tt.expectedTurns)
			}
			if !reflect.DeepEqual(distrib, tt.expectedDistrib) {
				t.Errorf("findOptimalTurns() distribution = %v, want %v", distrib, tt.expectedDistrib)
			}
		})
	}
}
