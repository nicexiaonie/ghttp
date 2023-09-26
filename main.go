package ghttp

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"unsafe"
)

// Post 发送post请求
func Post(url string, context FromValues, header map[string]string, timeout time.Duration) (Result, error) {

	result := Result{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(context.Encode()))
	if err != nil {
		return result, err
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

func PostRetry(url string, context FromValues, header map[string]string, timeout time.Duration, retry int) (Result, error) {
	for i := 1; i <= retry; i++ {
		post, err := Post(url, context, header, timeout)
		if err == nil {
			return post, err
		}
	}
	return Result{}, errors.New(fmt.Sprintf("request error. url:%s", url))
}

// PostJson 发送PostJson请求
func PostJson(url string, context FromValues, header map[string]string, timeout time.Duration) (Result, error) {

	result := Result{}

	reader := bytes.NewReader(context.Encode())
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

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

func PostJsonRetry(url string, context FromValues, header map[string]string, timeout time.Duration, retry int) (Result, error) {
	for i := 1; i <= retry; i++ {
		post, err := PostJson(url, context, header, timeout)
		if err == nil {
			return post, err
		}
	}
	return Result{}, errors.New(fmt.Sprintf("request error. url:%s", url))
}

// Get 发送Get请求
func Get(url string, context FromValues, header map[string]string, timeout time.Duration) (Result, error) {

	result := Result{}
	url += "?" + string(context.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	client.Timeout = timeout

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

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

func GetRetry(url string, context FromValues, header map[string]string, timeout time.Duration, retry int) (Result, error) {
	for i := 1; i <= retry; i++ {
		post, err := Get(url, context, header, timeout)
		if err == nil {
			return post, err
		}
	}
	return Result{}, errors.New(fmt.Sprintf("request error. url:%s", url))
}
