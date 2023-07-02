package snake

import (
	"log"
	"math/rand"

	"github.com/olivermking/wasmamba/model"
)

const (
	apiVersion = "1"
	author     = "OliverMKing"
	color      = "#161F24" // black

	// customizations from https://play.battlesnake.com/customizations
	head = "tongue"
	tail = "block-bum"
)

type snake struct{}

func New() *snake {
	return &snake{}
}

func (s *snake) Info() *model.InfoResp {
	return &model.InfoResp{
		ApiVersion: apiVersion,
		Author:     author,
		Color:      color,
		Head:       head,
		Tail:       tail,
	}
}

func (s *snake) Move(m model.GameReq) *model.MoveResp {
	id := m.You.Id

	moves := validMoves(id, model.GameReqSet(m))
	if len(moves) == 0 {
		log.Print("No valid moves")
		return &model.MoveResp{
			Move: model.Down,
		}
	}

	move := moves[rand.Intn(len(moves))]
	return &model.MoveResp{
		Move: move,
	}
}
