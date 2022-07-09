package main

import (
	"log"
	"net/http"

	"github.com/To-Do-REST-API/controller"
	"github.com/To-Do-REST-API/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	println("Server is listnening at port:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
