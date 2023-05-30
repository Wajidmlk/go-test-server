package main

// We will import most of packages for each features like printing & routers
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// function is written to react when `/` route is hit.
// Currently only printing demo messages on terminal, but can send JSON resps. etc
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}