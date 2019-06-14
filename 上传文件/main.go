package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

// 最大上传大小
const maxUploadSize = 2 * 1024 * 1024 // 2 mb
// 上传路径
const uploadPath = "./tmp"

// 主函数
func main() {
	http.HandleFunc("/upload", uploadFileHandler())

	// 创建文件服务
	fs := http.FileServer(http.Dir(uploadPath))
	// 上传时截取掉files前缀
	http.Handle("/files/", http.StripPrefix("/files", fs))
	
	log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 上传文件的方法
func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 通过内置方法只读取文件的前半部分  如果出错，函数会报错，没有超过设定大小不会报错
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		// 验证大小
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// 获取上传文件的类型
		fileType := r.PostFormValue("type")
		// 获取上传文件
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		// 别忘了关掉文件
		defer file.Close()
		// 读取文件内容
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// 检验文件类型
		filetype := http.DetectContentType(fileBytes)
		switch filetype {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		// 创建要存储的文件名
		fileName := randToken(12)
		// 获取文件后缀
		fileEndings, err := mime.ExtensionsByType(fileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
		// 拼接完整的文件路径
		newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
		fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

		// 将文件写入到新路径下
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		// 关掉写入句柄
		defer newFile.Close() 
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("SUCCESS"))
	})
}

// 渲染错误
func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
// 生成指定长度的文件名
func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
