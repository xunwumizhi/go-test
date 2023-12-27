package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestHmac(t *testing.T) {
	raw := `{"name": "Tom", "age": 1}`
	encode(raw)
}

func encode(data string) {
	key := ""
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	cryptoData := hex.EncodeToString(h.Sum(nil))
	fmt.Println(cryptoData)
}

func TestBase64(t *testing.T) {
	jsonStr := ``
	encodeStr := base64.StdEncoding.EncodeToString([]byte(jsonStr))
	fmt.Println(encodeStr)
}

// AESEncrypt AES加密
func AESEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AESDecrypt AES解密
func AESDecrypt(crypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])

	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// PKCS5填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5去填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func TestAES(t *testing.T) {
	key := []byte("1234567812345678")
	origData := []byte("hello world")
	crypted, err := AESEncrypt(origData, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	hexEnStr := hex.EncodeToString(crypted)
	fmt.Println("加密后hexStr: ", hexEnStr)

	hexBs, _ := hex.DecodeString(hexEnStr)

	decrypted, err := AESDecrypt(hexBs, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("解密后：", string(decrypted))
}

func TestSha256(t *testing.T) {
	str := strconv.Itoa(int(time.Now().UnixNano()))
	bs := sha256.Sum256([]byte(str))
	res := hex.EncodeToString(bs[:])
	fmt.Println(res, len(res))
}
