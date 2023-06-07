package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
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
