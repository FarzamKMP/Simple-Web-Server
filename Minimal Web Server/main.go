package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler_defualt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome dear user! you can add /hello or /bye in URL")
}
func handler_hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello dear user!")
}
func handler_bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dont leave me :(")
}

func main() {

	fmt.Println("The server is up now ! at :")
	time := time.Now()
	fmt.Println(time)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler_hello)
	mux.HandleFunc("/bye", handler_bye)
	mux.HandleFunc("/", handler_defualt)
	http.ListenAndServe(":8080", mux)
}
