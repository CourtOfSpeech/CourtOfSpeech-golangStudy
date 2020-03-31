package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

//自动生成的测试
// func Test_errWrapper(t *testing.T) {
// 	type args struct {
// 		handler appHandler
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want func(writer http.ResponseWriter, request *http.Request)
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := errWrapper(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("errWrapper() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

//实现了 userError 接口
func (e testingUserError) Message() string {
	return string(e)
}

//自定义Error
func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

//IsNotExist
func errIsNotExist(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

//IsPermission
func errIsPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

//unknow
func errUnknow(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknow server")
}

//nil
func errNil(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errIsNotExist, 404, "Not Found"},
	{errIsPermission, 403, "Forbidden"},
	{errUnknow, 500, "Internal Server Error"},
	{errNil, 200, ""},
}

//根据学习视频写的http测试
func Test_errWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)

		f(response, request)

		//读返回的信息
		// b, _ := ioutil.ReadAll(response.Body)
		// //strings.Trim()去掉不需要的字符 ，这里是去除换行
		// body := strings.Trim(string(b), "\n")
		// if response.Code != tt.code || body != tt.message {
		// 	t.Errorf("expect (%d , %s); got (%d, %s)", tt.code, tt.message, response.Code, body)
		// }
		verifyResponse(response.Result(), t, tt.code, tt.message)

	}
}

//开启一个http server 服务来测试

func Test_errWrapper_inServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)

		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		//读返回的信息
		verifyResponse(resp, t, tt.code, tt.message)
	}
}

//读返回的信息
func verifyResponse(resp *http.Response, t *testing.T, expectCode int, expectMessage string) {
	b, _ := ioutil.ReadAll(resp.Body)
	//strings.Trim()去掉不需要的字符 ，这里是去除换行
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectCode || body != expectMessage {
		t.Errorf("expect (%d , %s); got (%d, %s)", expectCode, expectMessage, resp.StatusCode, body)
	}
}
