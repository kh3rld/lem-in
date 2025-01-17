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