// 핸들러를 등록 관리하는 파일
package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
	// Get UserInfo by /users/{id}
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "User ID:", vars["id"])
}

// making a new my handler
func NewHttpHandler() http.Handler {
	// 인스턴스를 만들고 해당 인스턴스에 등록해서 사용하는 예제 코드.
	mux := mux.NewRouter() // gorilla mux 사용법
	// mux := http.NewServeMux() //gorilla mux 로 대체
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)
	// mux.Handle("/foo", &fooHandler{})

	return mux
}
