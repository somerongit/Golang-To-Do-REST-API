package controller

import (
	"encoding/json"
	"net/http"

	"github.com/To-Do-REST-API/views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("Request received...")
		println(r.Method, "(", r.Proto, ") =>", r.RequestURI)
		if r.Method == http.MethodGet {
			data := views.Respose{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
