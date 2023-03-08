package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", router)
}
