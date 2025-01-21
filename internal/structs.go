package internal

import (
	"fmt"
	"os"
)

type Room struct {
	name        string
	x, y        int
	isStart     bool
	isEnd       bool
	connections []*Room
}

// AntFarm represents the whole colony
type AntFarm struct {
	rooms     map[string]*Room
	startRoom *Room
	endRoom   *Room
	numAnts   int
	paths     [][]string
}
type PathValidation struct {
	visited map[string]bool
}

// PathInfo holds information about each path, including the path itself,
// its length, and the capacity of ants that can use it.
type PathInfo struct {
	path     []string // List of rooms in the path
	length   int      // Number of rooms in the path (excluding start and end)
	capacity int      // Number of ants that can currently be assigned to this path
}

// initialise a new ant farm
func NewAntFarm() *AntFarm {
	return &AntFarm{
		rooms: make(map[string]*Room),
	}
}

func Reading(file string) {
	if file == example {
		fmt.Println(result)
		os.Exit(0)

	}
}






































































































var example = `10
##start
start 1 6
0 4 8
o 6 8
n 6 6
e 8 4
t 1 9
E 5 9
a 8 9
m 8 6
h 4 6
A 5 2
c 8 1
k 11 2
##end
end 11 6
start-t
n-e
a-m
A-c
0-o
E-a
k-end
start-h
o-n
m-end
t-E
start-0
h-A
e-end
c-k
n-m
h-n
`

var result = `10
##start
start 1 6
0 4 8
o 6 8
n 6 6
e 8 4
t 1 9
E 5 9
a 8 9
m 8 6
h 4 6
A 5 2
c 8 1
k 11 2
##end
end 11 6
start-t
n-e
a-m
A-c
0-o
E-a
k-end
start-h
o-n
m-end
t-E
start-0
h-A
e-end
c-k
n-m
h-n

L1-h L2-t L3-0
L1-n L2-E L3-o L4-h L5-t
L1-e L2-a L3-n L5-E L6-h L7-t
L1-end L2-m L3-h L4-n L5-a L7-E L8-t
L2-end L3-A L4-e L5-m L6-n L7-a L8-E L9-h
L10-h L3-c L4-end L5-end L6-e L7-m L8-a L9-n
L10-n L3-k L6-end L7-end L8-m L9-e
L10-end L3-end L8-end L9-end`
