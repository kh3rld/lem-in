package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func (af *AntFarm) ParseRoom(line string, isStart bool, isEnd bool) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("ERROR: invalid data format, invalid room definition")
	}

	name := parts[0]
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return fmt.Errorf("ERROR: invalid data format, invalid room name")
	}

	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format, invalid x coordinate")
	}

	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format, invalid y coordinate")
	}

	room := &Room{
		name:        name,
		x:           x,
		y:           y,
		isStart:     isStart,
		isEnd:       isEnd,
		connections: make([]*Room, 0),
	}

	if isStart {
		if af.startRoom != nil {
			return fmt.Errorf("ERROR: invalid data format, multiple start rooms")
		}
		af.startRoom = room
	}
	if isEnd {
		if af.endRoom != nil {
			return fmt.Errorf("ERROR: invalid data format, multiple end rooms")
		}
		af.endRoom = room
	}

	af.rooms[name] = room
	return nil
}
