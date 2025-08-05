package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type:", "text/plain; charset=utf-8")
		body := []byte("OK")
		w.Write(body)
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.Handle("/assets", http.FileServer(http.Dir("assets/logo.png")))
	server := http.Server{}
	server.Handler = mux
	server.Addr = ":8080"
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
