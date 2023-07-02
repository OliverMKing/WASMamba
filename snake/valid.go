package snake

import "github.com/olivermking/wasmamba/model"

type moves func(id string, g model.GameReqS) []model.Move

func validMoves(id string, g model.GameReqS) []model.Move {
	ret := []model.Move{}
	opts := model.GetMoves()

	invalid := invalidSpaces(g)
	for _, opt := range opts {
		move := opt.NextCoord(g.You.Head)
		if _, ok := invalid[move]; !ok {
			if inBounds(move, g.Board.Height, g.Board.Width) {
				ret = append(ret, opt)
			}
		}
	}

	return ret
}

var _ moves = validMoves

type spaces func(g model.GameReqS) model.CoordS

func invalidSpaces(g model.GameReqS) model.CoordS {
	ret := make(model.CoordS)

	youHealth := g.You.Health
	hazardDmg := g.Game.Rules.Settings.HazardDamagePerTurn
	if youHealth-hazardDmg <= 0 {
		for c := range g.Board.Hazards {
			ret[c] = struct{}{}
		}
	}

	for _, s := range g.Board.Snakes {
		ret[s.Head] = struct{}{}

		// TODO: handle tail moving position on the next move
		for c := range s.Body {
			ret[c] = struct{}{}
		}
	}

	return ret
}

var _ spaces = invalidSpaces

func inBounds(c model.Coord, height, width int) bool {
	if c.X >= width || c.X < 0 {
		return false
	}

	if c.Y >= height || c.Y < 0 {
		return false
	}

	return true
}
