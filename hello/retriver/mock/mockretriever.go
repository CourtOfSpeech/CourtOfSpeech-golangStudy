package mock

import "fmt"

type Retriever1 struct {
	Contents string
}

//实现一个接口，只需要实现他的Get方法即可
func (r *Retriever1) Get(url string) string {
	return r.Contents
}

//Post
func (r *Retriever1) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

//常用的系统接口 go语言自带的 这里是 String 相当于java里的tostring
//其他常用的还有 Reader/Writer
func (r *Retriever1) String() string {
	return fmt.Sprintf(
		"Retriever: {Contents = %s}", r.Contents)
}
