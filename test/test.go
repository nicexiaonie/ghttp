package main

import (
	"fmt"
	"github.com/nicexiaonie/ghttp"
	"time"
)

func main() {
	requestParams := ghttp.FromValues{}
	requestParams.Add("name", "1111")
	requestParams.Add("phone", "1111")
	requestParams.Add("password", "1111")
	req, err := ghttp.PostJsonRetry("https://saishi.hainanxingdong.com/manage/club/create", requestParams, map[string]string{}, time.Second*6, 1)

	fmt.Printf("err0r:  %s", err)
	fmt.Printf("%+v", req)

}
