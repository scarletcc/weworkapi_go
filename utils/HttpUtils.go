package utils

import (
	"strings"
	"github.com/scarletcc/weworkapi_go/config"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
)

const base = "https://qyapi.weixin.qq.com"

func MakeUrl(queryArgs string) string {
	if strings.Index(queryArgs, "/") == 0 {
		return base + queryArgs
	}
	return base + "/" + queryArgs
}

func HttpGet(url string) ([]byte, error) {
	if config.DEBUG {
		fmt.Println("httpGet: " + url)
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func HttpPost(url, postData string) ([]byte, error) {
	if config.DEBUG {
		fmt.Println("httpPost: " + url)
		fmt.Println("postData: " + postData)
	}
	res, err := http.Post(url, "application/json", strings.NewReader(postData))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func HttpFile(url string, file io.Reader) ([]byte, error) {
	if config.DEBUG {
		fmt.Println("httpFile: " + url)
	}
	res, err := http.Post(url, "multipart/form-data", file)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
