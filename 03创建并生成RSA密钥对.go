/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:学习并掌握非对称加密中公钥和私钥的生成方法
 * @Date:2019/6/23 19:22
 * @Version:v1.0
*/
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

/**
  * @Description:创建公钥私钥密钥对，位数由自己指定，位数越大，安全性越高，效率越低。
  * @Params：参数bits指定生成密钥的长度
  */
func generateRsaKeyPair(bits int) error {
	fmt.Println("生成私钥")
	/*使用GenerateKey函数生成私钥，格式：
	 *func GenerateKey(random io.Reader, bits int) (priv *PrivateKey, err error)
	 *GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	 */
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Println("Generate Key err:", err)
		return err
	}
	/*对私钥进行编码，生成der格式的字符串
	 *x509包：公钥标准, func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte
	 */
	derStream, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Println("Marshal Key err:", err)
		return err
	}
	/*将der字符串拼接pem的编码结构*/
	block := pem.Block{
		Type:    "RSA Private Key",
		Headers: nil,
		Bytes:   derStream,
	}
	f1, err := os.Create("rsaPriKey.pem")
	if err != nil {
		return err
	}
	defer f1.Close()
	/*pem格式进行base64编码，得到最终的私钥*/
	err = pem.Encode(f1, &block)
	if err != nil {
		return nil
	}

	fmt.Println("生成公钥")
	/*通过私钥得到公钥*/
	publicKey := privateKey.PublicKey
	/*对公钥格式进行编码，生辰der格式的字符串*/
	pub, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	/*将der格式的字符串拼接到pem格式的block块中*/
	block = pem.Block{
		Type:    "RSA Public Key",
		Headers: nil,
		Bytes:   pub,
	}
	/*对pem格式的数据进行base64编码，得到最终格式的数据*/
	f2, err := os.Create("rsaPublicKey.pem")
	if err != nil {
		return nil
	}
	defer f1.Close()
	err = pem.Encode(f2, &block)
	if err != nil {
		return err
	}
	return nil
}

func main(){
	err := generateRsaKeyPair(1024)
	if err != nil {
		log.Println(err)
	}

}