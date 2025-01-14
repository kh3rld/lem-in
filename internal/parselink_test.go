package internal

import (
	"strings"
	"testing"
)

func TestParseLink(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		setup   func() *AntFarm
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid link",
			line: "start-room1",
			setup: func() *AntFarm {
				af := NewAntFarm()
				af.rooms = map[string]*Room{
					"start": {name: "start"},
					"room1": {name: "room1"},
				}
				return af
			},
			wantErr: false,
		},
		{
			name: "Self linking room",
			line: "room1-room1",
			setup: func() *AntFarm {
				af := NewAntFarm()
				af.rooms = map[string]*Room{
					"room1": {name: "room1"},
				}
				return af
			},
			wantErr: true,
			errMsg:  "ERROR: invalid data format, room cannot link to itself",
		},
		{
			name: "Invalid link format",
			line: "room1-room2-room3",
			setup: func() *AntFarm {
				af := NewAntFarm()
				return af
			},
			wantErr: true,
			errMsg:  "ERROR: invalid data format, invalid link format",
		},
		{
			name: "Link to unknown room",
			line: "start-unknown",
			setup: func() *AntFarm {
				af := NewAntFarm()
				af.rooms = map[string]*Room{
					"start": {name: "start"},
				}
				return af
			},
			wantErr: true,
			errMsg:  "ERROR: invalid data format, link to unknown room",
		},
		{
			name: "Duplicate link",
			line: "start-room1",
			setup: func() *AntFarm {
				af := NewAntFarm()
				room1 := &Room{name: "room1"}
				start := &Room{name: "start"}
				af.rooms = map[string]*Room{
					"start": start,
					"room1": room1,
				}
				start.connections = append(start.connections, room1)
				room1.connections = append(room1.connections, start)
				return af
			},
			wantErr: true,
			errMsg:  "ERROR: invalid data format, duplicate link",
		},
		{
			name: "Complex valid link chain",
			line: "room2-room3",
			setup: func() *AntFarm {
				af := NewAntFarm()
				af.rooms = map[string]*Room{
					"start": {name: "start"},
					"room1": {name: "room1"},
					"room2": {name: "room2"},
					"room3": {name: "room3"},
					"end":   {name: "end"},
				}
				return af
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := tt.setup()
			err := af.Parselink(tt.line)

			// Check if error matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("Parselink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If expecting specific error message, verify it
			if tt.wantErr && err != nil && err.Error() != tt.errMsg {
				t.Errorf("Parselink() error message = %v, want %v", err.Error(), tt.errMsg)
			}

			// For valid cases, verify the connections were properly created
			if !tt.wantErr {
				parts := strings.Split(tt.line, "-")
				room1 := af.rooms[parts[0]]
				room2 := af.rooms[parts[1]]

				// Check if bidirectional connection exists
				found1 := false
				for _, conn := range room1.connections {
					if conn == room2 {
						found1 = true
						break
					}
				}

				found2 := false
				for _, conn := range room2.connections {
					if conn == room1 {
						found2 = true
						break
					}
				}

				if !found1 || !found2 {
					t.Errorf("Parselink() failed to create bidirectional connection between %s and %s", parts[0], parts[1])
				}
			}
		})
	}
}
