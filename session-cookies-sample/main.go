package main

import (
	"fmt"
	"net/http"
	"session-cookies-sample/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/logout", handler.LogoutHandler)

	fmt.Println("Server is running on: 8080...")
	_ = http.ListenAndServe(":8080", mux)

}
