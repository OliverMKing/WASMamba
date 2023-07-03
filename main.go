package main

import (
	"encoding/json"
	"log"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/julienschmidt/httprouter"
	"github.com/olivermking/wasmamba/model"
	"github.com/olivermking/wasmamba/snake"
)

const (
	contentType           = "Content-Type"
	applicationJson       = "application/json"
	internalServiceErrMsg = "internal service error"
)

type Snake interface {
	Info() *model.InfoResp
	Move(model.GameReq) *model.MoveResp
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

func info(s Snake) spinhttp.RouterHandle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := json.NewEncoder(w).Encode(s.Info()); err != nil {
			http.Error(w, internalServiceErrMsg, http.StatusInternalServerError)
			log.Printf("failed to encode: %s", err)
			return
		}
		w.Header().Set(contentType, applicationJson)
	}
}

func move(s Snake) spinhttp.RouterHandle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var req model.GameReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(w).Encode(s.Move(req)); err != nil {
			http.Error(w, internalServiceErrMsg, http.StatusInternalServerError)
			log.Printf("failed to encode: %s", err)
			return
		}
		w.Header().Set(contentType, applicationJson)
	}
}

func noOp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

func main() {}
