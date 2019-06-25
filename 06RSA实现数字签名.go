/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:了解并使用RSA数字签名的方法
 * @Date:2019/6/24 12:32
 * @Version:v1.0
*/
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)
/**
 *私钥签名，公钥解签名，私钥签名时签名的是内容的哈希值
 */
func rsaSignedData(filename string, src []byte) ([]byte, error) {
	priKey, err := ReadRSAPriKey(filename)
	if err != nil {
		return nil, err
	}
	//求解hash值
	hash := sha256.Sum256(src)
	/*
	数字签名
	func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error) {
	得到签名数据
	参数1：随机数
	参数2: 私钥
	参数3: 计算哈希方法
	参数4: 原文的哈希值
	 */
	rsaSignedData, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	return rsaSignedData, nil
}

func rsaVerifyData(filename string, src []byte, rsaSigned []byte) (bool, error) {
	pubKey, err := ReadRSAPublicKey(filename)
	if err != nil {
		return false, err
	}
	//求解hash值
	hash := sha256.Sum256(src)
	/*
	签名认证
	func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error {
	参数1: 公钥
	参数2: 哈希算法
	参数3: 本地对原文生成的哈希
	参数4: 待验证数字签名
	*/
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], rsaSigned)
	if err != nil {
		return false, err
	}
	return true, nil
}
func main() {
	src := []byte("//公钥，需要校验的数据，数字签名（3个）")
	signData, err := rsaSignedData("rsaPriKey.pem", src)
	if err != nil {
		fmt.Println("sigData err:", err)
		return
	}

	fmt.Printf("signData : %x\n", signData)
	src = []byte("//公钥，需要校验的数据，数字签名（3个）11")

	b, _ := rsaVerifyData("rsaPublicKey.pem", src, signData)
	fmt.Println("res:", b)
}
