package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Greet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
}

func main() {
	mx := mux.NewRouter()

	mx.HandleFunc("/", SayHelloWorld)
	mx.HandleFunc("/{name}", Greet)

	http.ListenAndServe(":8080", mx)
}
