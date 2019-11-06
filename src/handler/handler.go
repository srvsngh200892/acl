package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
    "strconv"
	"github.com/gorilla/mux"
	"github.com/srvsngh200892/acl/src/role"
	"github.com/srvsngh200892/acl/src/user"
)

func CreateRoles(w http.ResponseWriter, r *http.Request) {
	var requestRoles []role.Role
	if err := json.NewDecoder(r.Body).Decode(&requestRoles); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST")
		return
	}

	if err := role.SetRoles(requestRoles); err !=nil {
		sendErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var requestUsers []user.User
	if err := json.NewDecoder(r.Body).Decode(&requestUsers); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST")
		return
	}
	if err := user.SetUsers(requestUsers); err !=nil {
		sendErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ListSubOrdinates(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path) 
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "BAD_REQUEST")
		fmt.Fprintln(w, err)
		return
	}

	subordinates, err := user.GetSubOrdinates(int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(subordinates)
}

// HealthCheck endpoint called to check if the app is uo
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "OK")
}

// Homepage handler redirects
func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.deputy.com", http.StatusMovedPermanently)
}

func sendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})

}
