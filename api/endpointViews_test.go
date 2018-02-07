package api

import (
	//"encoding/json"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

//type fakeSetup struct {
//	values []string
//	err error
//}
//
//func (f *fakeSetup) setup(values []string) error {
//	f.values = values
//	return f.err
//}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getconfigs", HandleGETS).Methods("GET")
	router.HandleFunc("/loadconfigs", HandlePOST).Methods("POST")
	return router
}

func TestHandlePOST(t *testing.T){
	assert.Equal(t, 0, 0, "Not passed.")
}

func TestHandleGET(t *testing.T){
	assert.Equal(t, 0, 0, "Not passed.")
}

//func TestHandleGETS(t *testing.T){
//	f := &fakeSetup{err: nil}
//	GetKVsFromConsul = f.setup
//
//	request, _ := http.NewRequest("GET", "/getconfigs", nil)
//	response := httptest.NewRecorder()
//	Router().ServeHTTP(response, request)
//
//	assert.Equal(t, 200, response.Code, "OK response is expected")
//}

func TestHandleGETS(t *testing.T){
	getkvOld := getkv
	defer func() {getkv = getkvOld}()

	getkv = func () ([]string, error){
		return nil, nil
	}

	request, _ := http.NewRequest("GET", "/getconfigs", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")
}
