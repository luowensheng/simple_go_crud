package main

import (
	"fmt"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	image "github.com/luowensheng/backend/api/image"
	text "github.com/luowensheng/backend/api/text"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Welcome to GOrilla")
	})

	router.HandleFunc("/api/v1/text", text.TextHandler)    //.Methods("POST", "GET")
	router.HandleFunc("/api/v1/image", image.ImageHandler) //.Methods("POST", "GET")

	router.HandleFunc("/api/v1/text/{id}", text.TextHandler)    //.Methods("POST", "GET")
	router.HandleFunc("/api/v1/image/{id}", image.ImageHandler) //.Methods("POST", "GET")

	fmt.Println("Listening on http://localhost:8080/ ")
	log.Fatal(http.ListenAndServe(":8080", router))
}
