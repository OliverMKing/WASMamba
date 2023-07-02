package snake

import "github.com/olivermking/wasmamba/model"

type path func(id string) (model.Move, error)
