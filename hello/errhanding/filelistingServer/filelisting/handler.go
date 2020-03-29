package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

//处理请求业务逻辑的方法
func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	//获取请求文件所在的路径 http://localhost:9999/list/fib.txt
	path := request.URL.Path[len("/list/"):]
	//获取文件后打开文件
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		//处理错误，将错误返回给页面  这样不能统一处理，并且把错误直接给用户是不好的体验
		// http.Error(writer,
		// 	err.Error(),
		// 	http.StatusInternalServerError)
		//报错后 要return
		return err
	}
	//打开文件后一定要及时关闭资源
	defer file.Close()

	//读文件
	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	//将文件写入ResponseWriter
	writer.Write(all)

	//如果没有报错
	return nil
}
