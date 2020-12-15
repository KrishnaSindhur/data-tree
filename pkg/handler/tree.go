package handler

import (
	"encoding/json"
	"net/http"

	"github.com/KrishnaSindhur/data-tree/pkg/contract"
	"github.com/KrishnaSindhur/data-tree/pkg/lib"
	"github.com/KrishnaSindhur/data-tree/pkg/service"

	"github.com/rs/zerolog/log"
)

var Tree = contract.Node{}

func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestQuery contract.Query
		if err := json.NewDecoder(r.Body).Decode(&requestQuery); err != nil {
			log.Error().Msg("Malformed request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := service.GetData(requestQuery, Tree)
		if err != nil {
			log.Err(err)
		}

		w.WriteHeader(http.StatusOK)
		lib.WriteResponseJSON(w, data)
	}
}

func Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestData contract.Data
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			log.Error().Msg("Malformed request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := service.AddData(requestData, &Tree)
		if err != nil {
			log.Err(err)
		}
		w.WriteHeader(http.StatusOK)
	}
}
