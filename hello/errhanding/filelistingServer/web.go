package main

import (
	"hello/errhanding/filelistingServer/filelisting"
	"log"
	"net/http"
	"os"
)

//错误统一处理
type appHandler func(writer http.ResponseWriter, request *http.Request) error

//处理错误
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			//输出日志
			log.Printf("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			//os.IsNotExist 判断文件是否存在
			case os.IsNotExist(err):
				code = http.StatusNotFound
			//os.IsPermission判断是否有权限
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				//StatusInternalServerError= 500 表示什么都不知道，不知道错误的类型
				code = http.StatusInternalServerError
			}
			//返回错误
			//http.StatusText 返回字符串文本
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

//写一个服务器，用来查看文件内容
func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	//开启一个服务器
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
