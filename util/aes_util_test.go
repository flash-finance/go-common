package util

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAAA(t *testing.T) {
	var aeskey = []byte("7894abu9y8d2fwop")
	pass := []byte("hello world")
	xpass, err := AesEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
}
