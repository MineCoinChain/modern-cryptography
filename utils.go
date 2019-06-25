package main

import (
	"crypto/rsa"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"github.com/pkg/errors"
)
//读取公钥文件转换为公钥
func ReadRSAPublicKey(filename string) (*rsa.PublicKey, error) {
	//读取公钥文件
	info, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	//解码,得到block
	block, _ := pem.Decode(info)
	//得到der字符串
	der := block.Bytes
	//得到公钥
	pubInterface, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, err
	}
	pubKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Generate PublicKey failed")
	}
	return pubKey, nil
}
//读取私钥文件转换为私钥
func ReadRSAPriKey(filename string) (*rsa.PrivateKey, error) {
	//读取私钥文件
	info, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	//解码，得到block
	block, _ := pem.Decode(info)
	//得到der
	der := block.Bytes

	//得到私钥
	priKeyInter, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}
	priKey, ok := priKeyInter.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("priKey no ok!")
	}
	return priKey, nil
}
