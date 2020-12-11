package lib

import (
	"encoding/json"
	"io"

	"github.com/KrishnaSindhur/data-tree/pkg/contract"

	"github.com/rs/zerolog/log"
)

func WriteResponseJSON(w io.Writer, responseData interface{}) {
	response := contract.Response{Output: responseData}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error().Err(err).Msg("Could not write JSON response")
	}
}
