package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	src := []byte("base64编码不是加解密，一种编辑数据的格式，方便传输。")
	//标准格式
	encodeInfo := base64.StdEncoding.EncodeToString(src)
	fmt.Println("encode info:%v\n", encodeInfo)
	//URL格式
	urlencodeInfo := base64.URLEncoding.EncodeToString(src)
	fmt.Print("URL encode info:%v\n", urlencodeInfo)
}
