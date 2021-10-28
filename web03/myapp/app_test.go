package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
설치방법

1. go get github.com/smartystreets/goconvey
PS D:\workspace\GO\tuckersGo\goWeb> go get github.com/smartystreets/goconvey
go: downloading github.com/smartystreets/goconvey v1.7.2
go: downloading golang.org/x/tools v0.0.0-20190328211700-ab21143f2384
go get: installing executables with 'go get' in module mode is deprecated.
        Use 'go install pkg@version' instead.
        For more information, see https://golang.org/doc/go-get-install-deprecation
        or run 'go help get' or 'go help install'.
PS D:\workspace\GO\tuckersGo\goWeb>

위에꺼 설치하고 path에 추가해줘야 함 ===> "%GOPATH%\bin"

2. go get github.com/stretchr/testify
D:\workspace\GO\tuckersGo\goWeb\web03>go get github.com/stretchr/testify
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading github.com/davecgh/go-spew v1.1.0
go: downloading github.com/stretchr/objx v0.1.0
go: downloading gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
go get: added github.com/davecgh/go-spew v1.1.0
go get: added github.com/pmezard/go-difflib v1.0.0
go get: added github.com/stretchr/objx v0.1.0
go get: added github.com/stretchr/testify v1.7.0
go get: added gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c

// 2번째에서 전체를 가져와서 아래 명령어 실행시 다운로드가 없음.
3. go get github.com/stretchr/testify/assert
D:\workspace\GO\tuckersGo\goWeb\web03>go get github.com/stretchr/testify/assert

D:\workspace\GO\tuckersGo\goWeb\web03>

*/
func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// barHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))

}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// barHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello name: World!", string(data))

}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=kimbs", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// barHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello name: kimbs!", string(data))

}
