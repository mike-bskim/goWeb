package model

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// 인터페이스 외부로 공개를 해야함, close 함수의 권한을 main.go 에 넘겨주기위해서
// 인터페이스 이름 및 내부 함수들을 대문자로변경, 외부에 공개
type DBHandler interface {
	// private 처리함
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

// init -> NewDBHandler
func NewDBHandler(filepath string) DBHandler {
	// handler := newMemoryHandler()
	// sqlite hanlder 만들면 아래처럼 변경만 하면 됨.
	// handler := newSqliteHandler()
	return newSqliteHandler(filepath)
}
