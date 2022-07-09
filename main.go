package main

import (
	"net/http"

	"github.com/To-Do-REST-API/controller"
)

func main() {
	mux := controller.Register()
	println("Server is listnening at port:3000")

	http.ListenAndServe(":3000", mux)
}
