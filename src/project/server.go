package project

import (
	"fmt"
	"net/http"
	"time"
)

func NewServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/digimovie", func(w http.ResponseWriter, r *http.Request) {fmt.Fprintf(w, "enter digimovie")})
	mux.HandleFunc("/digimovie/set/movie", Set_Movie)
	mux.HandleFunc("/digimovie/get/movie", Get_Movie)
	mux.HandleFunc("/digimovie/movies/list", Movies_List)
	mux.HandleFunc("/digimovie/version", Show_Version)

	server1 := http.Server{
		Addr: ":1386",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler: mux,
	}
	err := server1.ListenAndServe()
	if err != nil {
		panic(err)
	}
}