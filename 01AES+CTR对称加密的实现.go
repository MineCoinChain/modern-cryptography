/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:Encryption
 * @Description:了解并掌握AES-ATR加密解密的实现
 * @Date:2019/6/23 11:21
 * @Version:v1.0
*/
package main

import (
"bytes"
"crypto/aes"
"crypto/cipher"
"fmt"
"log"
)

/*
 *	AES的特点：密钥长度16,分组长度16
 *	CTR的特点：不需要进行填充,需要提供数据
*/

/*	加密算法的实现
 *	palintext: 需要进行加密的明文
 * 	key：加密使用到的密钥
 *  返回加密后的数据以及错误信息*/
func aesCtrEncryption(plaintext []byte, key []byte) ([]byte, error) {
	/*创建aes密码接口，使用aes包下的函数创建一个cipher.Block函数*/
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	/*创建分组模式CTR，采用crypto/cipher包，func NewCTR(block Block, iv []byte) Stream，
	*其中block为分组接口，iv为初始向量
	*iv要与算法的长度一致，16字节，这里使用bytes.Repeat()实现
	*/
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)
	/*使用XORKeyStream(dst, src []byte)进行加密，dst是密文空间，src是明文空间*/
	dst := make([]byte, len(plaintext))
	stream.XORKeyStream(dst, plaintext)
	return dst, nil
}

/*	解密算法的实现 */
func aesCtrDecryption(encryptData []byte, key []byte) ([]byte, error) {
	return aesCtrEncryption(encryptData, key)
}
func main() {
	//设置需要进行加密的明文
	plaintext := []byte("Hello World")
	//设置对称密钥
	key := []byte("1234567887654321")
	//调用加密函数进行加密，并输出加密后的结果
	encryptData, err := aesCtrEncryption(plaintext, key)
	if err != nil {
		log.Fatal("加密过程出现错误：", err)
	}
	fmt.Printf("encryptData is %X\n", encryptData)
	//调用解密函数进行解密，并输出解密后的结果
	plaintext, err = aesCtrDecryption(encryptData, key)
	if err != nil {
		log.Fatal("解密过程出现错误：", err)
	}
	fmt.Printf("decryptData is %s\n", string(plaintext))
}
