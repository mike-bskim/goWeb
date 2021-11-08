/*
PS D:\workspace\GO\tuckersGo\goWeb\web12> go mod init GO/tuckersGo/goWeb/web12
go: creating new go.mod: module GO/tuckersGo/goWeb/web12
go: to add module requirements and sums:
        go mod tidy
PS D:\workspace\GO\tuckersGo\goWeb\web12> go mod tidy
PS D:\workspace\GO\tuckersGo\goWeb\web12> go get -u github.com/gorilla/mux
go get: added github.com/gorilla/mux v1.8.0
PS D:\workspace\GO\tuckersGo\goWeb\web12>
고릴라보다 가벼운 라우팅 패키지
PS D:\workspace\GO\tuckersGo\goWeb\web12> go get github.com/gorilla/pat
go: downloading github.com/gorilla/pat v1.0.1
go: downloading github.com/gorilla/context v1.1.1
go get: added github.com/gorilla/context v1.1.1
go get: added github.com/gorilla/pat v1.0.1
PS D:\workspace\GO\tuckersGo\goWeb\web12>
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/pat"
)

type User struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "kimbs", Email: "kimbs@kimbs.com"}
	user.CreateAt = time.Now()

	w.Header().Add("Content-type", "application/json:")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreateAt = time.Now()
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func main() {
	// mux := http.NewServeMux() // basic mux
	// mux := mux.NewRouter() // gorila mux
	// mux.HandleFunc("/users", getUserInfoHandler).Methods("GET") 	// gorila mux 인 경우
	// mux.HandleFunc("/users", addUserHandler).Methods("PUT") 		// gorila mux 인 경우

	mux := pat.New() // gorila pat
	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)

	http.ListenAndServe(":3000", mux)
}
