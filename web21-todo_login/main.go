package main

import (
	"GO/tuckersGo/goWeb/web21-todo_login/myapp"
	"log"
	"net/http"
)

const portNumber = ":3000"

func main() {
	mux := myapp.MakeNewHandler("./todo.db")
	defer mux.Close()

	log.Println("Started App")
	err := http.ListenAndServe(portNumber, mux)
	if err != nil {
		panic(err)
	}
}
