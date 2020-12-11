package handler

import (
	"context"
	"net/http"
)

type TreeReader interface {
	GetData(ctx context.Context) (error)
}

func GetData(treeReader TreeReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
	}
}
