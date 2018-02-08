package api

import (
	//"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getconfigs", HandleGETS).Methods("GET")
	router.HandleFunc("/loadconfigs", HandlePOST).Methods("POST")
	return router
}

func TestHandlePOST(t *testing.T) {
	// TODO(sshank)
	assert.Equal(t, 0, 0, "Not passed.")
}

func TestHandleGET(t *testing.T) {
	// TODO(sshank)
	assert.Equal(t, 0, 0, "Not passed.")
}

func TestHandleGETS(t *testing.T) {
	getkvOld := getkvs
	defer func() { getkvs = getkvOld }()

	getkvs = func() ([]string, error) {
		return nil, nil
	}

	request, _ := http.NewRequest("GET", "/getconfigs", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")
}
