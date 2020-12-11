package handler

import (
	"encoding/json"
	"net/http"

	"github.com/KrishnaSindhur/data-tree/pkg/contract"
	"github.com/KrishnaSindhur/data-tree/pkg/lib"
	"github.com/KrishnaSindhur/data-tree/pkg/service"

	"github.com/rs/zerolog/log"
)

var tree = contract.Tree{}

func Get(tree contract.Tree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestQuery contract.Query
		if err := json.NewDecoder(r.Body).Decode(&requestQuery); err != nil {
			log.Error().Msg("Malformed request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := service.GetData(r.Context(), requestQuery, tree)
		if err != nil {
			log.Err(err)
		}

		w.WriteHeader(http.StatusOK)
		lib.WriteResponseJSON(w, data)
	}
}

func Add(tree contract.Tree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestData contract.Data
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			log.Error().Msg("Malformed request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		updatedTree, err := service.AddData(r.Context(), requestData, tree)
		if err != nil {
			log.Err(err)
		}
		tree = updatedTree
		w.WriteHeader(http.StatusOK)
	}
}
