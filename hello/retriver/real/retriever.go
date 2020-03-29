package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

//Retriever的实现...
type Retriever2 struct {
	UserAgent string
	TimeOut   time.Duration
}

//实现一个接口，只需要实现他的方法即可
func (r *Retriever2) Get(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}

	return string(result)
}
