package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil{
		// エラーハンドリング
		fmt.Println("エラー:", err)
		return
	}
	lw := logWriter{}

	io.Copy(lw,resp.Body)
}

func (logWriter) Write(bs []byte) (int,error) {
	fmt.Println(string(bs))
	fmt.Println("書き込んだバイト数",len(bs))
	return len(bs),nil
}
