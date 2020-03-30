package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/lists/"

type userErr string

func (e userErr) Error() string {
	return e.Message()
}

//实现了 userError 接口
func (e userErr) Message() string {
	return string(e)
}

//处理请求业务逻辑的方法
func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	//判断是否有这个方法处理逻辑的请求路径 ,没有就要返回
	if strings.Index(request.URL.Path, prefix) != 0 {
		//返回到go 自带的错误处理
		//return errors.New("path must start with " + prefix)

		//返回到自定义错误处理
		return userErr("path must start with " + prefix)
	}
	//获取请求文件所在的路径 http://localhost:9999/list/fib.txt
	path := request.URL.Path[len(prefix):]
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
