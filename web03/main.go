package main

import (
	"GO/tuckersGo/goWeb/myapp"
	"net/http"
)

const portNumber = ":3000"

func main() {

	mux := myapp.NewHttpHandler()
	http.ListenAndServe(portNumber, mux)
}
