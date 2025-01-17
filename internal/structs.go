package internal

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
