package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpAddr = ":8080"

func main()  {
	fmt.Println("Server running on",httpAddr)

	mux:= http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/standard-library-rocks", stdLibHandler)

	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	//_ := json.NewDecoder(r.Body).Decode(&)
	_, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func stdLibHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("I invite you to get curious about The Go programming language!"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
