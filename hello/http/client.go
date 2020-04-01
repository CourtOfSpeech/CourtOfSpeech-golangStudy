package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	url := "http://www.imooc.com"
	//http.NewRequest()能new一个request请求对象出来，这个对象能设置请求头等其他信息
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	//查看是否有重定向
	client := http.Client{
		CheckRedirect: func(
			//每次重定向的放req
			req *http.Request,
			//所有重定向的放在 via
			via []*http.Request) error {

			fmt.Println("Redirect :", req)
			//error == nil ,则会重定向，否则不会
			return nil
		},
	}

	//如果用对象就用下面这种方式发请求
	//resp, err := http.DefaultClient.Do(request)
	resp, err := client.Do(request)

	//不设置请求头
	//resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	//需要关闭
	defer resp.Body.Close()
	//解析resp	DumpResponse(resp *http.Response, body bool) ([]byte, error)
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
	}
	//打印返回的body
	fmt.Printf("%s\n", s)
}
