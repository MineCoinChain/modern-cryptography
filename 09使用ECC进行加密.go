
/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:test
 * @Description:学习并掌握go语言中椭圆曲线的用法
 * @Date:2019/6/30 21:50
 * @Version:v1.0
*/
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

/*go语言中只提供了椭圆曲线的签名，并未提供加解密*/

func main() {

	//需要进行签名的数据
	data := "hello world"
	//生成需要曲线
	curve := elliptic.P256()

	//创建私钥,参数:曲线，随机数
	priKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Fatal("generate key err;", err)
	}
	//获得公钥
	pubKey := priKey.PublicKey
	//对需要加密的内容做哈希
	hash := sha256.Sum256([]byte(data))
	//私钥签名
	r, s, err := ecdsa.Sign(rand.Reader, priKey, hash[:])
	if err != nil {
		log.Fatal("ecdsa sign err:", err)
	}
	//公钥验证
	//flag := ecdsa.Verify(&pubKey, hash[:], r, s)
	//通常不会直接像上边这种方式传输，而是将r和s转换成字节流传输,r和s转换成字节流后均为32
	//拼接r和s
	signature:=append(r.Bytes(),s.Bytes()...)

	//传输过程.....

	//对字节流进行转换
	r1:=signature[:len(signature)/2]
	s1:=signature[len(signature)/2:]
	var r2,s2 big.Int

	flag := ecdsa.Verify(&pubKey, hash[:], r2.SetBytes(r1), s2.SetBytes(s1))

	fmt.Println("flag is",flag)
}
