/*
 * @modifiy by:Mine&Coin&Chain
 * @Filename:main
 * @Description:学习并掌握des-cbc加密解密的实现
 * @Date:2019/6/23 12:07
 * @Version:v1.0
*/

package main

import (
"bytes"
"crypto/cipher"
"crypto/des"
"fmt"
"log"
)
/*
 *	ES的特点：密钥长度8,分组长度8
 *	CBC的特点：长度与算法相同，需要进行填充
*/

/*	加密算法的实现
 *	palintext: 需要进行加密的明文
 * 	key：加密使用到的密钥
 *  返回加密后的数据以及错误信息*/
func desCbcEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	/*创建es密码接口，使用des包下的函数创建一个cipher.Block函数*/
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	/*创建分组模式CBC，采用crypto/cipher包，func NewCBCEncrypter(block Block, iv []byte) BlockMode，
	*其中block为分组接口，iv为初始向量
	*返回一个密码分组链接模式的、底层用b解密的BlockMode接口
	*/
	var iv = bytes.Repeat([]byte("1"), block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	/*调用填充函数，对数据进行填充*/
	plaintext, err = paddingNumber(plaintext, block.BlockSize())
	if err != nil {
		return nil, err
	}
	/*调用加密函数
	type BlockMode interface {
		// 返回加密字节块的大小
		BlockSize() int
		// 加密或解密连续的数据块，src的尺寸必须是块大小的整数倍，src和dst可指向同一内存地址
		CryptBlocks(dst, src []byte)
	}
	*/
	mode.CryptBlocks(plaintext, plaintext)
	return plaintext, nil
}
/*加密时用到的填充函数*/
func paddingNumber(src []byte, blocksize int) ([]byte, error) {
	/*
		* leftNum；分组之后的剩余长度
		* needNum：得到需要填充的个数
	*/
	leftNum := len(src) % blocksize
	needNum := blocksize - leftNum
	newSlice := bytes.Repeat([]byte{byte(needNum)}, needNum)
	src = append(src, newSlice...)
	return src, nil
}
/**
	* 解密算法的实现
	* 解密算法实现过程和加密算法相似，不同之处在于解密算法需要先解密，再去填充。
 */
func desCbcDecrypt(encryptData []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var iv = bytes.Repeat([]byte("1"), block.BlockSize())
	mode := cipher.NewCBCDecrypter(block, iv)
	//对数据先进行解密
	mode.CryptBlocks(encryptData, encryptData)
	//对解密后的数据去除填充
	encryptData = unpaddingNum(encryptData)
	return encryptData, nil
}
/*去除填充的实现函数*/
func unpaddingNum(src []byte) []byte {
	leftNum := src[len(src)-1]
	return src[:len(src)-int(leftNum)]
}
func main() {
	text := "hello world"
	key := "12345678"
	//加密处理
	encryptData, err := desCbcEncrypt([]byte(text), []byte(key))
	if err != nil {
		log.Fatal("Ercrypt err:", err)
	}
	fmt.Printf("Encrypt Data: %x\n", encryptData)
	//解密处理
	plaintext, err := desCbcDecrypt(encryptData, []byte(key))
	if err != nil {
		log.Fatal("Decrypt err:", err)
	}
	fmt.Printf("Decrypt Data: %s\n", plaintext)
}
