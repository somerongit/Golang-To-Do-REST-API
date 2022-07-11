package main

import (
	"log"
	"net/http"
	"os"

	"github.com/To-Do-REST-API/controller"
	"github.com/To-Do-REST-API/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	port := os.Getenv("PORT")
	println("Server is listnening at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
