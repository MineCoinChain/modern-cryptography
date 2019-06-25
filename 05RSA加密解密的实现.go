package main

import (
	"crypto/rsa"
	"crypto/rand"
	"fmt"
)

//公钥加密
func rsaEncryptData(filename string, src []byte) ([]byte, error) {
	//获取公钥
	pubkey, err := ReadRSAPublicKey(filename)
	if err != nil {
		return nil, err
	}
	//用公钥进行加密
	encryptInfo, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, src)
	if err != nil {
		return nil, err
	}
	return encryptInfo, err
}

//私钥解密
func rsaDecryptData(filename string, src []byte) ([]byte, error) {
	priKey, err := ReadRSAPriKey(filename)
	if err != nil {
		return nil, err
	}
	//用私钥进行解密
	info, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, src)
	if err != nil {
		return nil, err
	}
	return info, nil
}
func main() {
	src := []byte("hello world")
	info, err := rsaEncryptData("/home/itcast/workspace/go/src/rsaPublicKey.pem", src)
	if err != nil {
		fmt.Println("加密过程报错:", err)
		return
	}

	fmt.Println("加密后的数为:", info)

	plainText, err := rsaDecryptData("/home/itcast/workspace/go/src/rsaPriKey.pem", info)
	if err != nil {
		fmt.Println("解密过程报错:", err)
		return
	}
	fmt.Printf("明文数据:%s\n", plainText)
}
