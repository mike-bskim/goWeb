package main

import (
	"GO/tuckersGo/goWeb/web18-todo_sqlite/myapp"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

const portNumber = ":3000"

func main() {
	mux := myapp.MakeNewHandler()
	defer mux.Close()

	ng := negroni.Classic()
	ng.UseHandler(mux)

	log.Println("Started App")
	err := http.ListenAndServe(portNumber, ng)
	if err != nil {
		panic(err)
	}

}
