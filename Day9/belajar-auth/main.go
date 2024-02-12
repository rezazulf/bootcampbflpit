package main

import (
	middleware "belajar-auth"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", middleware.ActionStudent)

	server := new(http.Server)
	server.Addr = ":3000"

	fmt.Println("Server started at port 3000")
	server.ListenAndServe()
}
