package main

import (
	"fmt"
	"net/http"
	"time"
)

const portNumber = ":3000"

type User struct {
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// user := new(User)
	// err := json.NewDecoder(r.Body).Decode(user)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(w, "Bad Request: ", err)
	// 	return
	// }
	fmt.Fprintf(w, "Hello Foo!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello name: %s!", name)
	// fmt.Fprintf(w, "Hello Bar!")
}

func main() {
	// 인스턴스를 만들고 해당 인스턴스에 등록해서 사용하는 예제 코드.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(portNumber, mux)
}
