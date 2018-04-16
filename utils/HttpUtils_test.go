package utils

import (
	"testing"
	"fmt"
	_ "encoding/json"
)

func TestMakeUrl(t *testing.T) {
	if MakeUrl("/aaa/bb?t=1") != "https://qyapi.weixin.qq.com/aaa/bb?t=1" {
		t.Error("test1 /aaa/bb?t=1 fail")
	}
	if MakeUrl("aaa/bb?t=1") != "https://qyapi.weixin.qq.com/aaa/bb?t=1" {
		t.Error("test2 aaa/bb?t=1 fail")
	}
}

func TestHttpGet(t *testing.T) {
	body, err := HttpGet("https://www.baidu.com/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

