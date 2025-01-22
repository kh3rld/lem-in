package internal

import (
	"strings"
	"testing"
)

func TestAntFarm_ParseRoom(t *testing.T) {
	type fields struct {
		rooms     map[string]*Room
		startRoom *Room
		endRoom   *Room
		numAnts   int
		paths     [][]string
	}

	type args struct {
		line    string
		isStart bool
		isEnd   bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		errMsg  string
	}{
		// Valid room creation
		{
			name: "Valid Normal Room",
			fields: fields{
				rooms: make(map[string]*Room),
			},
			args: args{
				line:    "room1 10 20",
				isStart: false,
				isEnd:   false,
			},
			wantErr: false,
		},
		// Start room creation
		{
			name: "Valid Start Room",
			fields: fields{
				rooms: make(map[string]*Room),
			},
			args: args{
				line:    "start 0 0",
				isStart: true,
				isEnd:   false,
			},
			wantErr: false,
		},
		// End room creation
		{
			name: "Valid End Room",
			fields: fields{
				rooms: make(map[string]*Room),
			},
			args: args{
				line:    "end 100 100",
				isStart: false,
				isEnd:   true,
			},
			wantErr: false,
		},
		// Invalid room name starting with L
		{
			name: "Invalid Room Name (L prefix)",
			fields: fields{
				rooms: make(map[string]*Room),
			},
			args: args{
				line:    "Lroom 10 20",
				isStart: false,
				isEnd:   false,
			},
			wantErr: true,
			errMsg:  "invalid room name",
		},
		// Invalid room name starting with #
		{
			name: "Invalid Room Name (# prefix)",
			fields: fields{
				rooms: make(map[string]*Room),
			},
			args: args{
				line:    "#room 10 20",
				isStart: false,
				isEnd:   false,
			},
			wantErr: true,
			errMsg:  "invalid room name",
		},
		// Invalid coordinate (non-numeric)
		{
			name: "Invalid X Coordinate",
			fields: fields{
				rooms: make(map[string]*Room),
			},
			args: args{
				line:    "room1 ten 20",
				isStart: false,
				isEnd:   false,
			},
			wantErr: true,
			errMsg:  "invalid x coordinate",
		},
		// Invalid coordinate (non-numeric)
		{
			name: "Invalid Y Coordinate",
			fields: fields{
				rooms: make(map[string]*Room),
			},
			args: args{
				line:    "room1 10 twenty",
				isStart: false,
				isEnd:   false,
			},
			wantErr: true,
			errMsg:  "invalid y coordinate",
		},
		// Multiple start rooms
		{
			name: "Multiple Start Rooms",
			fields: fields{
				rooms:     make(map[string]*Room),
				startRoom: &Room{name: "existing_start"},
			},
			args: args{
				line:    "new_start 0 0",
				isStart: true,
				isEnd:   false,
			},
			wantErr: true,
			errMsg:  "multiple start rooms",
		},
		// Multiple end rooms
		{
			name: "Multiple End Rooms",
			fields: fields{
				rooms:   make(map[string]*Room),
				endRoom: &Room{name: "existing_end"},
			},
			args: args{
				line:    "new_end 100 100",
				isStart: false,
				isEnd:   true,
			},
			wantErr: true,
			errMsg:  "multiple end rooms",
		},
		// Insufficient room data
		{
			name: "Insufficient Room Data",
			fields: fields{
				rooms: make(map[string]*Room),
			},
			args: args{
				line:    "room1 10",
				isStart: false,
				isEnd:   false,
			},
			wantErr: true,
			errMsg:  "invalid data format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AntFarm{
				rooms:     tt.fields.rooms,
				startRoom: tt.fields.startRoom,
				endRoom:   tt.fields.endRoom,
				numAnts:   tt.fields.numAnts,
				paths:     tt.fields.paths,
			}
			err := af.ParseRoom(tt.args.line, tt.args.isStart, tt.args.isEnd)

			if (err != nil) != tt.wantErr {
				t.Errorf("AntFarm.ParseRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("Expected error message to contain %q, got %q", tt.errMsg, err.Error())
			}
		})
	}
}
