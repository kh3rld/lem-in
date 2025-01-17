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

func TestGenerateMoves(t *testing.T) {
	tests := []struct {
		name        string
		paths       []PathInfo
		turns       int
		numAnts     int
		endRoomName string
		expected    []string
	}{
		{
			name: "Single ant, single path",
			paths: []PathInfo{
				{
					path:     []string{"start", "room1", "end"},
					length:   2,
					capacity: 1,
				},
			},
			turns:       2,
			numAnts:     1,
			endRoomName: "end",
			expected:    []string{"L1-room1", "L1-end"},
		},
		{
			name: "Multiple ants, multiple paths",
			paths: []PathInfo{
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
			turns:       3,
			numAnts:     3,
			endRoomName: "end",
			expected: []string{
				"L1-room1 L2-room2",
				"L1-end L2-end L3-room1",
				"L3-end",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateMoves(tt.paths, tt.turns, tt.numAnts, tt.endRoomName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("generateMoves() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMoveExistingAnts(t *testing.T) {
	paths := []PathInfo{
		{
			path:     []string{"start", "room1", "end"},
			length:   2,
			capacity: 1,
		},
	}
	
	antStates := make(map[int]struct {
		pathIndex int
		position  int
	})
	antStates[1] = struct {
		pathIndex int
		position  int
	}{0, 1}
	
	occupied := make(map[string]bool)
	var currentMoves []string
	
	tests := []struct {
		name           string
		endRoomName    string
		expectedMoves  []string
		expectedStates map[int]struct {
			pathIndex int
			position  int
		}
	}{
		{
			name:        "Move ant to end",
			endRoomName: "end",
			expectedMoves: []string{"L1-end"},
			expectedStates: map[int]struct {
				pathIndex int
				position  int
			}{
				1: {0, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveExistingAnts(&antStates, paths, occupied, &currentMoves, tt.endRoomName)
			
			if !reflect.DeepEqual(currentMoves, tt.expectedMoves) {
				t.Errorf("moveExistingAnts() moves = %v, want %v", currentMoves, tt.expectedMoves)
			}
			if !reflect.DeepEqual(antStates, tt.expectedStates) {
				t.Errorf("moveExistingAnts() states = %v, want %v", antStates, tt.expectedStates)
			}
		})
	}
}