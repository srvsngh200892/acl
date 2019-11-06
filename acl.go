package main

import (
	"log"
	"net/http"
	"github.com/srvsngh200892/acl/src/router"
)

func main() {
	http.Handle("/", router.NewRouter())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
