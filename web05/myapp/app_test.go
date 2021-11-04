package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer((NewHttpHandler()))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer((NewHttpHandler()))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(string(data), "No Users")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer((NewHttpHandler()))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/1")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println("TestGetUserInfo:", string(data))
	assert.Contains(string(data), "No User ID:1")
}

/*
go get -u github.com/gorilla/mux
*/

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer((NewHttpHandler()))
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"bs", "last_name":"kim", "email":"kimbs@kimbs.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.Id)

	id := user.Id
	log.Println("TestCreateUser user.Id:", id)
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.Id, user2.Id)
	assert.Equal(user.FirstName, user2.FirstName)

}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer((NewHttpHandler()))
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"bs", "last_name":"kim", "email":"kimbs@kimbs.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, err := ioutil.ReadAll(resp.Body)
	log.Println("TestDeleteUser:", string(data))
	assert.NoError(err)
	assert.Equal("Deleted User ID:1", string(data))

}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer((NewHttpHandler()))
	defer ts.Close()

	// update wrong user
	req, _ := http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(`{"id":1, "first_name":"updated", "last_name":"updated", "email":"updated@kimbs.com"}`))
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// no user updated
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:1")
	log.Printf("TestUpdateUser:[%s]", string(data))

	// create test user info
	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"bs", "last_name":"kim", "email":"kimbs@kimbs.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.Id)

	// update the test user info from create
	newUser := new(User)
	newUser.Id = user.Id
	newUser.FirstName = "updated"
	newUser.LastName = "up"
	newUser.Email = "up@up.com"

	updateStr := fmt.Sprintf(`{"id":%d, "first_name":"%s", "last_name":"%s", "email":"%s"}`,
		newUser.Id, newUser.FirstName, newUser.LastName, newUser.Email)

	req, _ = http.NewRequest("PUT", ts.URL+"/users", strings.NewReader(updateStr))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// data, _ = ioutil.ReadAll(resp.Body)
	// log.Printf("user.Id[%d] resp:[%s]", user.Id, string(data))

	// checking update result from updateUserHandler
	updateUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(updateUser)
	assert.NoError(err)
	log.Printf("TestUpdateUser Id[%d] data:[%v]", updateUser.Id, *updateUser)

}

func TestUsers_withUsersData(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer((NewHttpHandler()))
	defer ts.Close()

	// create test user info
	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"bs", "last_name":"kim", "email":"kimbs@kimbs.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	// create test user info
	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"jason", "last_name":"park", "email":"jason@kimbs.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	// data, err := ioutil.ReadAll(resp.Body)
	// log.Printf("Users:[%s]", string(data))
	// assert.NoError(err)
	// assert.NotZero(len(data))
	users := []*User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(2, len(users))
	for _, d := range users {
		log.Print("TestUsers_withUsersData:", *d)
	}
}
