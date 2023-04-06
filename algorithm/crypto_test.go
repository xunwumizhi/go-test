package main

import (
	"crypto/hmac"
	"crypto/sha256"
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
