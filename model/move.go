package model

type Move string

const (
	Up    Move = "up"
	Down  Move = "down"
	Left  Move = "left"
	Right Move = "right"
)

func GetMoves() []Move {
	return []Move{Up, Down, Left, Right}
}

func (m Move) NextCoord(c *Coord) *Coord {
	switch m {
	case Up:
		return &Coord{
			X: c.X,
			Y: c.Y + 1,
		}
	case Down:
		return &Coord{
			X: c.X,
			Y: c.Y - 1,
		}
	case Left:
		return &Coord{
			X: c.X - 1,
			Y: c.Y,
		}
	case Right:
		return &Coord{
			X: c.X + 1,
			Y: c.Y,
		}
	}

	return nil
}
