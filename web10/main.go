package main

import (
	"GO/tuckersGo/goWeb/web10/decoHandler"
	"GO/tuckersGo/goWeb/web10/myapp"
	"log"
	"net/http"
	"time"
)

const portNumber = ":3000"

func logger1(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Printf("[LOGGER1] Started\n")
	h.ServeHTTP(w, r)
	log.Printf("[LOGGER1] Completed, time: %v ms\n", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Printf("[LOGGER2] Started\n")
	h.ServeHTTP(w, r)
	log.Printf("[LOGGER2] Completed, time: %v ms\n", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	mux := myapp.NewHttpHandler()
	mux = decoHandler.NewDecoHandler(mux, logger1)
	mux = decoHandler.NewDecoHandler(mux, logger2)
	return mux
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(portNumber, mux)
}
