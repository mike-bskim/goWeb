package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "C:\\Users\\MIKE-mini\\Downloads\\BlueStacks 002.png"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	// 웹으로 파일을 전송하는 포멧은 MIME 이걸만드려면 멀티파트가 필요함.
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	// multi , 폼파일을 만들었음.
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path)) // filepath.Base(path) 파일이름을 짤라줌.
	assert.NoError(err)
	// file -> multi 로 파일을 복사함. buf 로 데이터가 저장됨.
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)

	// 헤더에 콘텐츠의 타입을 알려줘야 함, form 타입으로 알려줌.
	req.Header.Set("Content-type", writer.FormDataContentType())

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{} // 바이트 어레이
	originData := []byte{} // 바이트 어레이
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)

}
