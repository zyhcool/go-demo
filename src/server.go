package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zyhcool/go-demo/morestrings"
	"github.com/zyhcool/go-demo/sayhello"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	w.Write([]byte("hello this is / response"))
	fmt.Println(morestrings.ReverseRunes("hello"))
	sayhello.SayIt()
}

func handlertest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("this is test")
	w.Write([]byte("hello this is test response"))
}

func main() {
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/test", handlertest)
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/test", handlertest)
	http.Handle("/", r)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("server start at 3000")
	}
}
