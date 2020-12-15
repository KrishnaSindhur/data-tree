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

func Test_Add_ShouldBeAbleToAddDataToTree(t *testing.T) {
	requestBody := `{"dim": [{"key": "device","val": "mobile"},{"key": "country","val": "IN"}],"metrics": [{"key": "webreq","val": 70},{"key": "timespent","val": 30}]}`
	req, _ := http.NewRequest(http.MethodPost, "/data-tree/v1/insert", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	MuxRouter(handler.Add(), "/data-tree/v1/insert", http.MethodPost).ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Incorrect HTTP status code")
}

func Test_Add_ShouldBeAbleToAddDataToTreeIfSomeKeyIsMissing(t *testing.T) {
	requestBody := `{"dim": [{"key": "country","val": "IN"}],"metrics": [{"key": "webreq","val": 70},{"key": "timespent","val": 30}]}`
	req, _ := http.NewRequest(http.MethodPost, "/data-tree/v1/insert", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	MuxRouter(handler.Add(), "/data-tree/v1/insert", http.MethodPost).ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Incorrect HTTP status code")
}

func Test_Get_ShouldBeAbleToGetDataFromTree(t *testing.T) {
	tree := contract.Node{}
	tree.WebReq = 140
	tree.TimeSpent = 60
	var node2 [] *contract.Node
	node2 = append(node2, &contract.Node{MetaData: "mobile", WebReq: 140, TimeSpent: 60})
	node1 := contract.Node{MetaData: "IN", WebReq: 140, TimeSpent: 60, Children: node2}
	tree.Children = append(tree.Children, &node1)
	handler.Tree = tree

	requestBody := `{"dim" : [{"key": "country","val": "IN"}]}`
	expectedValue := `{"output":{"dim":[{"key":"country","val":"IN"}],"metrics":[{"key":"webreq","val":140},{"key":"timespent","val":60}]}}
`
	req, _ := http.NewRequest(http.MethodGet, "/data-tree/v1/query", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	MuxRouter(handler.Get(), "/data-tree/v1/query", http.MethodGet).ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Incorrect HTTP status code")
	assert.Equal(t, expectedValue, w.Body.String(), "Incorrect response")
}

func MuxRouter(handler http.Handler, path string, method string) *mux.Router {
	if method == "" {
		method = http.MethodGet
	}
	r := mux.NewRouter()
	r.Handle(path, handler).Methods(method)
	return r
}
