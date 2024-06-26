package maps

import (
	"context"
	"errors"
	"math/rand"
)

// Build creates a map with the given size.
func Build(ctx context.Context, size, shipCount int) (*Map, error) {
	if size < 3 {
		return nil, errors.New("map size too small, min allowed size if 9 x 9")
	}

	if shipCount%2 != 0 {
		return nil, errors.New("each player must have the same number of ships")
	}

	m := &Map{size: size, board: make(board)}
	return addShip(m, Red)
}

func addShip(m *Map, color Color) (*Map, error) {
	ship := &Ship{
		size:        3,
		orientation: Orientation(rand.Int31() % 2),
		color:       color,
	}

	if err := m.addShip(ship); err != nil {
		return nil, err
	}

	return m, nil
}

// Map is a 2D representation of world where the game will be played.
type Map struct {
	size int

	ships []*Ship

	board board
}

type board map[int]map[int]bool

func (b board) setTrue(x, y int) {
	if b[x] == nil {
		b[x] = map[int]bool{}
	}
	b[x][y] = true
}

func (m *Map) addShip(s *Ship) error {
	var added bool
	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			if m.canBePlace(i, j, s) {
				for k := 0; k < s.size; k++ {
					if s.orientation == NorthSouth {
						m.board.setTrue(i+k, j)
					} else {
						m.board.setTrue(i, j+k)
					}
				}
				added = true
			}
		}
	}

	if !added {
		return errors.New("ship cannot be placed anywhere")
	}

	m.ships = append(m.ships, s)

	return nil
}

func (m *Map) canBePlace(x, y int, s *Ship) bool {
	return true
}

type Ship struct {
	playerID    int
	size        int
	orientation Orientation
	color       Color
}

type Orientation int

const (
	NorthSouth Orientation = iota
	WestEast
)

type Color string

const (
	Red  Color = "red"
	Blue Color = "blue"
)
