package internal

import (
	"os"
	"strings"
	"testing"
)

func createTempFile(t *testing.T, content string) string {
	tmpfile, err := os.CreateTemp("", "test-farm-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}
	return tmpfile.Name()
}

func TestAntFarm_ParseInput(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		wantErr     bool
		errMsg      string
	}{
		{
			name: "Valid farm configuration",
			fileContent: `4
##start
start 0 0
room1 1 1
room2 2 2
##end
end 3 3
start-room1
room1-room2
room2-end`,
			wantErr: false,
		},
		{
			name:        "Empty file",
			fileContent: "",
			wantErr:     true,
			errMsg:      "ERROR: invalid data format, empty file",
		},
		{
			name: "Invalid number of ants",
			fileContent: `0
##start
start 0 0
##end
end 1 1
start-end`,
			wantErr: true,
			errMsg:  "ERROR: invalid data format, invalid number of ants",
		},
		{
			name: "Missing start room",
			fileContent: `4
room1 1 1
##end
end 3 3
room1-end`,
			wantErr: true,
			errMsg:  "ERROR: invalid data format, no start room found",
		},
		{
			name: "Missing end room",
			fileContent: `4
##start
start 0 0
room1 1 1
start-room1`,
			wantErr: true,
			errMsg:  "ERROR: invalid data format, no end room found",
		},
		{
			name: "No path from start to end",
			fileContent: `4
##start
start 0 0
room1 1 1
##end
end 3 3
start-room1`,
			wantErr: true,
			errMsg:  "ERROR: invalid data format, no link from start to end",
		},
		{
			name: "Valid farm with comments",
			fileContent: `4
#comment
##start
start 0 0
#another comment
room1 1 1
room2 2 2
##end
end 3 3
start-room1
room1-room2
room2-end`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file with test content
			tmpfile := createTempFile(t, tt.fileContent)
			defer os.Remove(tmpfile)

			af := NewAntFarm()
			content, err := af.ParseInput(tmpfile)

			// Check error conditions
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseInput() expected error but got none")
				} else if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ParseInput() error = %v, wantErr = %v", err, tt.errMsg)
				}
				return
			}

			// Verify successful parsing
			if err != nil {
				t.Errorf("ParseInput() unexpected error: %v", err)
			}

			// Verify the content matches the input
			if !tt.wantErr && !strings.Contains(content, strings.TrimSpace(tt.fileContent)) {
				t.Errorf("ParseInput() content mismatch\ngot: %v\nwant: %v", content, tt.fileContent)
			}

			// Additional validations for successful cases
			if !tt.wantErr {
				if af.startRoom == nil {
					t.Error("ParseInput() start room not set")
				}
				if af.endRoom == nil {
					t.Error("ParseInput() end room not set")
				}
				if af.numAnts <= 0 {
					t.Error("ParseInput() invalid number of ants")
				}
			}
		})
	}
}

func TestAntFarm_ParseInput_FileErrors(t *testing.T) {
	af := NewAntFarm()

	// Test with non-existent file
	_, err := af.ParseInput("nonexistent.txt")
	if err == nil {
		t.Error("ParseInput() expected error for non-existent file")
	}

	// Test with directory instead of file
	tmpDir := t.TempDir()
	_, err = af.ParseInput(tmpDir)
	if err == nil {
		t.Error("ParseInput() expected error when passing directory")
	}
}
