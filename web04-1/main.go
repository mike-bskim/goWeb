package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const portNumber = ":3000"

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file") // id 랑 매칭.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename) // 파일이름은 헤더에 있음
	log.Println("filepath:", filepath)

	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func NewHttpHandler() http.Handler {
	// 인스턴스를 만들고 해당 인스턴스에 등록해서 사용하는 예제 코드.
	mux := http.NewServeMux()

	mux.HandleFunc("/uploads", uploadsHandler)
	mux.Handle("/", http.FileServer(http.Dir("public")))

	return mux
}

func main() {
	// http.HandleFunc("/uploads", uploadsHandler)
	// http.Handle("/", http.FileServer(http.Dir("public")))
	mux := NewHttpHandler()
	http.ListenAndServe(portNumber, mux)
}
