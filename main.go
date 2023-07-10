package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/julienschmidt/httprouter"
	"github.com/olivermking/wasmamba/logger"
	"github.com/olivermking/wasmamba/model"
	"github.com/olivermking/wasmamba/snake"
	"go.uber.org/zap"
)

const (
	contentType           = "Content-Type"
	applicationJson       = "application/json"
	internalServiceErrMsg = "internal service error"
)

type Snake interface {
	Info(context.Context) *model.InfoResp
	Move(context.Context, model.GameReq) *model.MoveResp
}

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		router := spinhttp.NewRouter()
		s := snake.New()

		// https://docs.battlesnake.com/api
		router.GET("/", info(s))
		router.POST("/start", noOp)
		router.POST("/move", move(s))
		router.POST("/end", noOp)

		router.ServeHTTP(w, r)
	})
}

func commonReq(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		start := time.Now()
		ctx := r.Context()
		ctx = logger.WithRequest(ctx, r)
		logger := logger.FromContext(ctx)
		defer logger.Sync()

		logger.Info("starting to handle request")
		defer func() {
			logger.Info("finished handling request", zap.String("responseTime", time.Since(start).String()))
		}()

		h(w, r.WithContext(ctx), p)
	}
}

func info(s Snake) spinhttp.RouterHandle {
	return commonReq(
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			ctx := r.Context()
			l := logger.FromContext(ctx)

			if err := json.NewEncoder(w).Encode(s.Info(ctx)); err != nil {
				l.Error(fmt.Sprintf("failed to encode response: %s", err.Error()))
				http.Error(w, internalServiceErrMsg, http.StatusInternalServerError)
				return
			}
			w.Header().Set(contentType, applicationJson)
		},
	)
}

func move(s Snake) spinhttp.RouterHandle {
	return commonReq(
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			ctx := r.Context()
			l := logger.FromContext(ctx)

			var game model.GameReq
			if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
				l.Error(fmt.Sprintf("failed to decode request: %s", err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			ctx = logger.WithGame(ctx, game)
			if err := json.NewEncoder(w).Encode(s.Move(ctx, game)); err != nil {
				l.Error(fmt.Sprintf("failed to encode response: %s", err.Error()))
				http.Error(w, internalServiceErrMsg, http.StatusInternalServerError)
				return
			}
			w.Header().Set(contentType, applicationJson)
		},
	)
}

func noOp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

func main() {}
