package main

import (
	"net/http"
)

func main() {
	server := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    server.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", server.getUsersHandler)
	mux.HandleFunc("POST /users", server.createUserHandler)

	srv.ListenAndServe()

}
