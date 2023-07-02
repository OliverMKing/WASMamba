package model

type Set[C comparable] map[C]struct{}

type CoordS Set[Coord]

// CoordSet converts a slice of Coords to a set
func CoordSet(c []Coord) CoordS {
	coords := make(CoordS)
	for _, coord := range c {
		coords[coord] = struct{}{}
	}
	return coords
}

type GameReqS struct {
	Game  Game   `json:"game"`
	Turn  int    `json:"turn"`
	Board BoardS `json:"board"`
	You   SnakeS `json:"you"`
}

func GameReqSet(g GameReq) GameReqS {
	board := BoardSet(g.Board)
	you := SnakeSet(g.You)
	return GameReqS{
		Game:  g.Game,
		Turn:  g.Turn,
		Board: board,
		You:   you,
	}
}

// BoardS is a board with coords in a set
type BoardS struct {
	Height  int      `json:"height"`
	Width   int      `json:"width"`
	Food    CoordS   `json:"food"`
	Hazards CoordS   `json:"hazards"`
	Snakes  []SnakeS `json:"snakes"`
}

func BoardSet(b Board) BoardS {
	food := CoordSet(b.Food)
	hazards := CoordSet(b.Hazards)

	snakes := []SnakeS{}
	for _, s := range b.Snakes {
		snakes = append(snakes, SnakeSet(s))
	}

	return BoardS{
		Height:  b.Height,
		Width:   b.Width,
		Food:    food,
		Hazards: hazards,
		Snakes:  snakes,
	}
}

// SnakeS is a snake with coords in a set
type SnakeS struct {
	Id             string        `json:"id"`
	Name           string        `json:"name"`
	Health         int           `json:"health"`
	Body           CoordS        `json:"body"`
	Head           Coord         `json:"head"`
	Length         int           `json:"length"`
	Latency        string        `json:"latency"`
	Shout          string        `json:"shout"`
	Customizations Customization `json:"customizations"`
}

func SnakeSet(s Snake) SnakeS {
	body := CoordSet(s.Body)
	return SnakeS{
		Id:             s.Id,
		Name:           s.Name,
		Health:         s.Health,
		Body:           body,
		Head:           s.Head,
		Length:         s.Length,
		Latency:        s.Latency,
		Shout:          s.Shout,
		Customizations: s.Customizations,
	}
}
