package server

// Import
import (
	"net/http"
	"log"
	"fmt"
)

// HTTP Handlers
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// Exported Server Functions
func Start() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Stop() {

}

func Reload() {

}