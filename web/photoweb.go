package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

const (
	UPLOAD_DIR = "E:\\document\\uploads"
)

/**
使用net/http包来一步步构建整个相册程序的网络服务<br>
上传
**/
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<html><body><form method=\"POST\" action=\"/upload\" enctype=\"multipart/form-data\"> "+
			"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
			"<input type=\"submit\" value=\"Upload\" /> </form></body></html>")
		return
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image") //寻找名为image的文件域并对其接值. 返回3个值:multipart.File/*multipart.FileHeader/error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

//查看指定文件是否存在
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

//查看上传的图片(单张查看)
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

//所有文件列表
func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var listHtml, li string
	for _, fileInfo := range fileInfoArr {
		imgid := fileInfo.Name()
		li += "<li><a href='/view?id=" + imgid + "'>" + imgid + "</a></li>"
	}
	listHtml = "<html><body><ol>" + li + "<ol></body></html>"
	io.WriteString(w, listHtml)
}

//闭包避免程序运行时出错崩溃: 所有handler都经过此方法,统一处理handler抛出的异常panic
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
func start_web() {
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/", safeHandler(listHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
