package myapp

import (
	"GO/tuckersGo/goWeb/web16-todo_refactoring/model"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(MakeNewHandler())
	defer ts.Close()

	// add testing - first data
	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	todo := new(model.Todo)
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo")
	id1 := todo.ID
	log.Println("app_test.go / add result >", *todo, id1)

	// add testing - second data
	resp, err = http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo2"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	// todo = new(Todo)
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo2")
	id2 := todo.ID
	log.Println("app_test.go / add result >", *todo, id2)

	// get testing - getting whole data
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	todos := []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)
	for i := 0; i < len(todos); i++ {
		log.Println("after getting whole data >", *todos[i])
	}

	// complete testing
	resp, err = http.Get(ts.URL + "/complete-todo/" + strconv.Itoa(id2) + "?complete=true")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// todos := []*Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)
	for i := 0; i < len(todos); i++ {
		log.Println("after complete for >", *todos[i])
	}

	// DELETE testing
	req, _ := http.NewRequest("DELETE", ts.URL+"/todos/"+strconv.Itoa(id1), nil) // data??? ??????????????? nil ??????
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	log.Println("delete id1: ", id1)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// todos := []*Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 1)
	for i := 0; i < len(todos); i++ {
		log.Println("after delete for >", *todos[i])
	}

}
