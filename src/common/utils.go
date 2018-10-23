package common

import (
	"net/http"
	"os"
	"strings"
)

//GetRequestURL 获取请求的URL
func GetRequestURL(request *http.Request) string {
	scheme := "http://"
	if request.TLS != nil {
		scheme = "https://"
	}

	return strings.Join([]string{scheme, request.Host, request.RequestURI}, "")
}

//PathExist 路径是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
