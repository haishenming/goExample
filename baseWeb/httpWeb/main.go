package main

import (
	"net/http"
	"fmt"
)



func main() {
	resp, err := http.Get("http://baidu.com/")

	fmt.Printf("%t, %t", resp, err)
}