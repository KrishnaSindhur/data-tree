package lib

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/KrishnaSindhur/data-tree/pkg/contract"
	"io"
)

func WriteResponseJSON(w io.Writer, responseData interface{}, logger zerolog.Logger) {
	response := contract.Response{Data: responseData}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Error().Err(err).Msg("Could not write JSON response")
	}
}
