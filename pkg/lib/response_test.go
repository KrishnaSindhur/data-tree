package lib_test

import (
	"bytes"
	"testing"

	"github.com/KrishnaSindhur/data-tree/pkg/lib"

	"github.com/stretchr/testify/assert"
)

func TestWriteResponseJSON(t *testing.T) {
	var w bytes.Buffer
	response := map[string]string{"A": "one", "B": "two"}

	lib.WriteResponseJSON(&w, response)

	assert.JSONEq(t, `{"output":{"A": "one", "B": "two"}}`, w.String(), "Incorrect response written")
}
