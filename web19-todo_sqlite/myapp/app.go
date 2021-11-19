package myapp

import (
	"GO/tuckersGo/goWeb/web18-todo_sqlite/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

// 인터페이스 외부로 공개를 해야함, close 함수의 권한을 main.go 에 넘겨주기위해서
// app.go 에도 언제 close 할지 모르므로 main.go 로 이관.
// 그래서 AppHandler 만들어서 넘겨줌.
type AppHandler struct {
	// 임베디드 처리함. 이유는 ??? 상속과 비슷한데 조금 다름 is 관계는 아니고, has 관계임.
	// 이름을 암시적으로 생략함. handler 생략함.
	http.Handler
	db model.DBHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {

	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := a.db.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

func (a *AppHandler) Close() {
	a.db.Close()
}

// 리턴변경 http.Handler -> AppHandler
func MakeNewHandler(filepath string) *AppHandler {

	mux := mux.NewRouter()
	a := &AppHandler{
		Handler: mux,
		db:      model.NewDBHandler(filepath),
	}
	mux.HandleFunc("/", a.indexHandler)
	mux.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	mux.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	return a
}
