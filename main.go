package ghttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
	"unsafe"
)

// 设置日志输出函数
func SetLogger(callFunc func(format string, args ...interface{})) {
	logger = loggerModel{
		Output: callFunc,
	}
}

// 设置http.Transport参数
func SetTransport(t http.Transport) {
	client.Transport = &t
}

// 发送post请求
func Post(requestUrl string, context []byte, header map[string]string, timeout time.Duration) (Result, error) {

	logger.Output("request url:%s, content:%s", requestUrl, context)
	result := Result{}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewReader(context))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client.Timeout = timeout

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	str := (*string)(unsafe.Pointer(&body))

	result.Body = *str
	result.StatusCode = resp.StatusCode
	result.Status = resp.Status
	result.Header = resp.Header
	result.ContentLength = resp.ContentLength

	return result, err

}

// 发送PostJson请求
func PostJson(url string, context []byte, header map[string]string, timeout time.Duration) (Result, error) {

	result := Result{}

	reader := bytes.NewReader(context)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return result, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client.Timeout = timeout
	resp, err := client.Do(request)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))

	result.Body = *str
	result.StatusCode = resp.StatusCode
	result.Status = resp.Status
	result.Header = resp.Header
	result.ContentLength = resp.ContentLength

	return result, err

}

// 发送Get请求
func Get(requestUrl string, context FromValues, header map[string]string, timeout time.Duration) (Result, error) {

	logger.Output("request url:%s, content:%s", requestUrl, context)
	result := Result{}
	requestUrl += "?" + context.Encode()
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return result, err
	}

	client.Timeout = timeout

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	str := (*string)(unsafe.Pointer(&body))

	result.Body = *str
	result.StatusCode = resp.StatusCode
	result.Status = resp.Status
	result.Header = resp.Header
	result.ContentLength = resp.ContentLength

	return result, err

}
