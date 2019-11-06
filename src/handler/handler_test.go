package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoles(t *testing.T) {

	payload := []byte(`[{"Id": 1, "Name": "System Administrator", "Parent": 0 },{ "Id": 2, "Name": "Location Manager", "Parent": 1 }, { "Id": 3, "Name": "Supervisor", "Parent": 2 }, { "Id": 4, "Name": "Employee", "Parent": 3 }, { "Id": 5, "Name": "Trainer", "Parent": 3 }]`)

	req, _ := http.NewRequest("POST", "/roles", bytes.NewBuffer(payload))
	response := executeRequest(req, CreateRoles)
    
    assert.Equal(t, http.StatusCreated, response.Code)
}

func TestInvalidCreateRoles(t *testing.T) {

	payload := []byte(`[{"Id": 1, "Parent": 0 },{ "Id": 2, "Name": "Location Manager", "Parent": 1 }, { "Id": 3, "Name": "Supervisor", "Parent": 2 }, { "Id": 4, "Name": "Employee", "Parent": 3 }, { "Id": 5, "Name": "Trainer", "Parent": 3 }]`)

	req, _ := http.NewRequest("POST", "/roles", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req, CreateRoles)
    
    assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestCreateUsers(t *testing.T) {

	payload := []byte(`[{"Id": 1, "Name": "Adam Admin", "Role": 1 }, { "Id": 2, "Name": "Emily Employee", "Role": 4 }, { "Id": 3, "Name": "Sam Supervisor", "Role": 3 }, { "Id": 4, "Name": "Mary Manager", "Role": 2 }, { "Id": 5, "Name": "Steve Trainer", "Role": 5 }]`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req, CreateUsers)
    
    assert.Equal(t, http.StatusCreated, response.Code)
}

func TestInvalidCreateUsers(t *testing.T) {

	payload := []byte(`{"Id": 0, "Name": "Adam Admin", "Role": 1 }, { "Id": 2, "Name": "Emily Employee", "Role": 4 }, { "Id": 3, "Name": "Sam Supervisor", "Role": 3 }, { "Id": 4, "Name": "Mary Manager", "Role": 2 }, { "Id": 5, "Name": "Steve Trainer", "Role": 5 }`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req, CreateUsers)
    
    assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestListSubOrdinates(t *testing.T) {
	req1, _ := http.NewRequest("GET", "/subordinates/3", nil)
	req1.Header.Set("Content-Type", "application/json")
	
	rr := httptest.NewRecorder()
	// We need a mux router in order to pass in the `name` variable.
	r := mux.NewRouter()

	r.HandleFunc("/subordinates/{id:.*}", ListSubOrdinates).Methods("GET")
	r.ServeHTTP(rr, req1)

    assert.Equal(t, http.StatusOK, rr.Code)

    expected := `[{"Id":2,"Name":"Emily Employee","Role":4},{"Id":5,"Name":"Steve Trainer","Role":5}]`
    actual := strings.TrimSuffix(rr.Body.String(), "\n")
    assert.Equal(t, actual, expected)
}

func TestEmptyListSubOrdinates(t *testing.T) {
	req1, _ := http.NewRequest("GET", "/subordinates/5", nil)
	req1.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	// We need a mux router in order to pass in the `name` variable.
	r := mux.NewRouter()

	r.HandleFunc("/subordinates/{id:.*}", ListSubOrdinates).Methods("GET")
	r.ServeHTTP(rr, req1)
    
    assert.Equal(t, http.StatusOK, rr.Code)
    actual := strings.TrimSuffix(rr.Body.String(), "\n")
    assert.Equal(t, actual, `[]`)
}

func executeRequest(req *http.Request, f func(http.ResponseWriter,
	*http.Request)) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(f)
	handler.ServeHTTP(rr, req)
	return rr
}
