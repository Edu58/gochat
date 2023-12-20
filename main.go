package main

import (
	"io"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	_, err := io.WriteString(w, `<h1>Hello World</h1>`)
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
