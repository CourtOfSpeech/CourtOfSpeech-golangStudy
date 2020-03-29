package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

//写一个服务器，用来查看文件内容
func main() {
	http.HandleFunc("/list/",
		func(writer http.ResponseWriter,
			request *http.Request) {
			//获取请求文件所在的路径
			path := request.URL.Path[len("/list/"):]
			//获取文件后打开文件
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			//打开文件后一定要及时关闭资源
			defer file.Close()

			//读文件
			all, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			//将文件写入ResponseWriter
			writer.Write(all)
		})

	//开启一个服务器
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
