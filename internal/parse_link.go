package internal

import (
	"fmt"
	"strings"
)

func (af *AntFarm) Parselink(line string) error {
	parts := strings.Split(line, "-")
	// Add in Parselink
	if parts[0] == parts[1] {
		return fmt.Errorf("ERROR: invalid data format, room cannot link to itself")
	}
	if len(parts) != 2 {
		return fmt.Errorf("ERROR: invalid data format, invalid link format")
	}

	room1, exists1 := af.rooms[parts[0]]
	room2, exists2 := af.rooms[parts[1]]

	if !exists1 || !exists2 {
		return fmt.Errorf("ERROR: invalid data format, link to unknown room")
	}

	for _, conn := range room1.connections {
		if conn == room2 {
			return fmt.Errorf("ERROR: invalid data format, duplicate link")
		}
	}

	room1.connections = append(room1.connections, room2)
	room2.connections = append(room2.connections, room1)
	return nil
}
