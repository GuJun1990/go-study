package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	fmt.Println(resp)
	fmt.Println(err)
}
