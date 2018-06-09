package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Gophercon 2017!")
	})

	port := ":8081"
	if v, ok := os.LookupEnv("PORT"); ok {
		port = ":" + v
	}
	log.Println("Listening on " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
