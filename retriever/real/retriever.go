package real

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r Retriever) Get (url string) string  {
	res, err := http.Get(url)
	if  err != nil{
		panic(err)
	}
	result, err := httputil.DumpResponse(res, true)
	res.Body.Close()
	if err != nil{
		panic(err)
	}
	return string(result)
}