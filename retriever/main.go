package main

import (
	"fmt"
	"learn_go_code/retriever/mock"
	"learn_go_code/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://blog.lychiyu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"mock data"}
	fmt.Printf("%T %v\n", r, r)
	fmt.Println(download(r))
	r = real.Retriever{}

	fmt.Printf("%T %v\n", r, r)
	fmt.Println(download(r))
}
