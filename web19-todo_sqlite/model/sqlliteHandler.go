package model

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos() []*Todo {
	todos := []*Todo{}
	sql_string := "SELECT id, name, completed, createdAt From todos"
	rows, err := s.db.Query(sql_string)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}
	return todos
}

func (s *sqliteHandler) AddTodo(name string) *Todo {
	sql_string := "INSERT INTO todos (name, completed, createdAt) VALUES(?,?,datetime('now'))"
	stmt, err := s.db.Prepare(sql_string)
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(name, false)
	if err != nil {
		panic(err)
	}

	id, _ := result.LastInsertId()
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()

	return &todo
}

func (s *sqliteHandler) RemoveTodo(id int) bool {
	sql_string := "DELETE FROM todos WHERE id = ?"
	stmt, err := s.db.Prepare(sql_string)
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := result.RowsAffected()

	return cnt > 0
}

func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	sql_string := "UPDATE todos SET completed = ? WHERE id = ?"
	stmt, err := s.db.Prepare(sql_string)
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(complete, id)
	if err != nil {
		panic(err)
	}
	cnt, _ := result.RowsAffected()

	return cnt > 0
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

func newSqliteHandler(filepath string) DBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id        INTEGER  PRIMARY KEY AUTOINCREMENT,
			name      TEXT,
			completed BOOLEAN,
			createdAt DATETIME
		)`)
	statement.Exec()
	return &sqliteHandler{db: database}
	// return &sqliteHandler{}
}
