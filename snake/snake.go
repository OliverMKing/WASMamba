package snake

import (
	"context"
	"math/rand"

	"github.com/olivermking/wasmamba/logger"
	"github.com/olivermking/wasmamba/model"
)

const (
	apiVersion = "1"
	author     = "OliverMKing"
	color      = "#74e4bc" // spin green

	// customizations from https://play.battlesnake.com/customizations
	head = "alligator"
	tail = "alligator"
)

type snake struct{}

func New() *snake {
	return &snake{}
}

func (s *snake) Info(_ context.Context) *model.InfoResp {
	return &model.InfoResp{
		ApiVersion: apiVersion,
		Author:     author,
		Color:      color,
		Head:       head,
		Tail:       tail,
	}
}

func (s *snake) Move(ctx context.Context, m model.GameReq) *model.MoveResp {
	logger := logger.FromContext(ctx)

	id := m.You.Id
	moves := validMoves(id, model.GameReqSet(m))
	if len(moves) == 0 {
		logger.Warn("no valid moves")
		return &model.MoveResp{
			Move: model.Down,
		}
	}

	move := moves[rand.Intn(len(moves))]
	return &model.MoveResp{
		Move: move,
	}
}
