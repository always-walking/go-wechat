package main

import (
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

func main() {
	info := common.CUserInfo{
		AppId:     "wxfedcab8946a21ccc",
		AppSecret: "7ae3cc7b23b34a7b8d8cea44f7e9177f",
		Port:      80,
		Url:       "/test/wechat/hello",
		Token:     "c0082b4d-b46f-4d67-b0eb-7a0d15dd5ef2",
	}
	wechat.New(&info)
}