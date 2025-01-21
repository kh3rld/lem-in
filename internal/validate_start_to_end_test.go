package internal

import "testing"

func TestAntFarm_ValidateStartEndPath(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "Valid farm with path",
			filename: "testfarms/validfarm.txt",
			wantErr:  false,
		},
		{
			name:     "Invalid farm - no path to end",
			filename: "testfarms/invalidfarm1.txt",
			wantErr:  true,
		},
		{
			name:     "Invalid farm - disconnected components",
			filename: "testfarms/invalidfarm2.txt",
			wantErr:  true,
		},
		{
			name:     "Invalid farm - circular path no end",
			filename: "testfarms/invalidfarm3.txt",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := NewAntFarm()
			_, err := af.ParseInput(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
