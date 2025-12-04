package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// コードとして長い。
	// bs := make([]byte, 99999) // create a byte slice(99999 は推定値であり、実際のデータサイズに応じて調整が必要)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}
