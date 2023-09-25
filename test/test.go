package main

import (
	"fmt"
	"github.com/nicexiaonie/ghttp"
	"time"
)

func main() {
	requestParams := ghttp.FromValues{}

	req, err := ghttp.GetRetry("https://www.baidu.com/", requestParams, map[string]string{}, time.Second*6, 10)

	fmt.Println(err)
	fmt.Printf("%+v", req)

}
