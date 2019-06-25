package main

import (
	"crypto/md5"
	"io"
	"fmt"
)

//go语言中调用hash函数的方法
func md51(src []byte) ([]byte, error) {
	//创建一个hasher接口
	hasher := md5.New()
	io.WriteString(hasher, string(src))
	hash := hasher.Sum([]byte("11"))
	return hash[:], nil
}

func md52(src []byte) ([]byte, error) {
	hash := md5.Sum(src)
	return hash[:], nil
}
func main() {
	src := []byte("hello world")
	hash, _ := md51(src)
	fmt.Println(hash)
	hash, _ = md52(src)
	fmt.Print(hash)

}
