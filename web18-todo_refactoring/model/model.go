package model

import (
	"log"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

func init() {
	// model.TodoMap = make(map[int]*model.Todo)
	todoMap = make(map[int]*Todo)
}

func GetTodos() []*Todo {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	return list
}

func AddTodo(name string) *Todo {
	id := len(todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo
	log.Println("add Todo success")
	// for i := 1; i < len(model.TodoMap)+1; i++ {
	// 	log.Println("app.go / addTodoHandler >", model.TodoMap[i])
	// }
	return todo
}

func RemoveTodo(id int) bool {

	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		// log.Println("delete success")
		// for i := 1; i < len(todoMap)+1; i++ {
		// 	log.Println(todoMap[i])
		// }
		// rd.JSON(w, http.StatusOK, Success{Success: true})
		return true
	}
	// else {
	// 	// log.Println("delete false", todoMap)
	// 	rd.JSON(w, http.StatusOK, Success{Success: false})
	// }
	return false
}

func CompleteTodo(id int, complete bool) bool {
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		return true
		// rd.JSON(w, http.StatusOK, Success{Success: true})
	}
	// else {
	// 	rd.JSON(w, http.StatusOK, Success{Success: false})
	// }

	return false
}
