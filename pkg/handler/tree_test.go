package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/KrishnaSindhur/data-tree/pkg/contract"
	"github.com/KrishnaSindhur/data-tree/pkg/handler"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

//Todo should mock the service layer
func Test_Add_ShouldBeAbleToPostData(t *testing.T) {
	requestBody := `{"dim": [{"key": "device","val": "mobile"},{"key": "country","val": "IN"}],"metrics": [{"key": "webreq","val": 70},{"key": "timespent","val": 30}]}`
	req, _ := http.NewRequest(http.MethodPost, "/data-tree/v1/insert", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	MuxRouter(handler.Add(contract.Tree{}), "/data-tree/v1/insert", http.MethodPost).ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Incorrect HTTP status code")
}

func Test_Add_ShouldBeAbleToPostDataIfSomeKeyIsMissing(t *testing.T) {
	requestBody := `{"dim": [{"key": "country","val": "IN"}],"metrics": [{"key": "webreq","val": 70},{"key": "timespent","val": 30}]}`
	req, _ := http.NewRequest(http.MethodPost, "/data-tree/v1/insert", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	MuxRouter(handler.Add(contract.Tree{}), "/data-tree/v1/insert", http.MethodPost).ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Incorrect HTTP status code")
}

func MuxRouter(handler http.Handler, path string, method string) *mux.Router {
	if method == "" {
		method = http.MethodGet
	}
	r := mux.NewRouter()
	r.Handle(path, handler).Methods(method)
	return r
}
