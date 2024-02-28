package grid

const GRID_SIZE int = 3

var directionVectorMap map[string]Coordinate = map[string]Coordinate{
	"left":  {0, -1},
	"right": {0, 1},
	"up":    {-1, 0},
	"down":  {1, 0},
}

type Coordinate [2]int

func Clamp(c int) int {
	return max(min(GRID_SIZE-1, c), 0)
}

func (c Coordinate) AddClamp(c2 Coordinate) Coordinate {
	return Coordinate{Clamp(c[0] + c2[0]), Clamp(c[1] + c2[1])}
}

func (c Coordinate) ToIdx() int {
  return (c[0] * GRID_SIZE) + (c[1]) + 1
}

func NextCoordInDirection(coord Coordinate, dir string) Coordinate {
  return coord.AddClamp(directionVectorMap[dir]) 
}

// Convert 1-based index to position in grid
func IdxToCoord(idx int) Coordinate {
	x := (idx - 1) / GRID_SIZE
	y := (idx - 1) % GRID_SIZE
	return Coordinate{x, y}
}
