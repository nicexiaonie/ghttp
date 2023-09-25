package ghttp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
)

var client *http.Client

func init() {
	tr := &http.Transport{
		MaxIdleConns:        5000,
		MaxIdleConnsPerHost: 2000,
	}
	client = &http.Client{
		Transport: tr,
	}
}

// Result 返回结果
type Result struct {
	StatusCode    int
	Status        string
	Body          string
	Header        http.Header
	ContentLength int64
}

// FromValues FormValues
type FromValues map[string]interface{}

func (v FromValues) EncodeJson() string {
	contextStr, _ := buildJson(v)
	return string(contextStr)
}

func (v FromValues) Encode() []byte {
	if v == nil {
		return []byte("")
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := v[k]
		prefix := url.QueryEscape(k) + "="
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(prefix)
		valueType := reflect.TypeOf(v).String()

		if valueType == "string" {
			value := v.(string)
			buf.WriteString(url.QueryEscape(value))
		} else if valueType == "int" {
			value := strconv.Itoa(v.(int))
			buf.WriteString(url.QueryEscape(value))
		} else if valueType == "int64" {
			value := strconv.FormatInt(v.(int64), 10)
			buf.WriteString(url.QueryEscape(value))
		} else if valueType == "uint64" {
			value := strconv.FormatUint(v.(uint64), 10)
			buf.WriteString(url.QueryEscape(value))
		} else if valueType == "http.FromValues" {
			value, _ := buildJson(v)
			buf.WriteString(string(value))
		}
	}
	return buf.Bytes()
}

func (v FromValues) Add(key string, value interface{}) {
	v[key] = value
}

func buildJson(data interface{}) ([]byte, error) {
	buf := bytes.NewBufferString("")
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(&data); err != nil {
		return nil, err
	} else {
		return buf.Bytes(), nil
	}
}
