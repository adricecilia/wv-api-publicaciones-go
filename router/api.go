package router

import (
	"fmt"
	"net/http"
)

func Init() {
	router()
	http.ListenAndServe(":8080", nil)
}

func router() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login, %s!", r.URL.Path[1:])
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home, %s!", r.URL.Path[1:])
}
