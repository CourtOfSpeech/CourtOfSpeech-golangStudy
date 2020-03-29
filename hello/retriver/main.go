package main

import (
	"fmt"
	"hello/retriver/mock"
	"hello/retriver/real"
	"time"
)

const url = "http://www.imooc.com"

// 定义一个接口，有一个Get方法
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

//组合接口
type RetrieverPoster interface {
	Retriever
	Poster
	//也可以定义自己的方法
	//Connect(host string)
}

func download(r Retriever) string {
	return r.Get(url)
}

func poster(p Poster) {
	p.Post(url, map[string]string{
		"name": "xinjiang",
		"age":  "18",
	})
}

func session(s RetrieverPoster) string {
	//s.Get(url)
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> Type: %T Value: %v\n", r, r)
	fmt.Print(" > Type switch: ")

	switch v := r.(type) {
	case *mock.Retriever1:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever2:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever

	mockRetriever := mock.Retriever1{
		Contents: "thsi is a fake imooc.com"}
	//接口在使用时一般不用指针，因为他本身包含实现者的指针
	r = &mockRetriever
	fmt.Println(download(r))
	inspect(r)
	r = &real.Retriever2{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Println(download(r))
	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever1); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&mockRetriever))

}
