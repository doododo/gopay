package wechat

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/go-pay/gopay/pkg/aes"
	"github.com/go-pay/gopay/pkg/util"
)

// 敏感信息加密，默认 PKCS1
func V3EncryptText(text string, rsaPublicKeyContent []byte) (cipherText string, err error) {
	block, _ := pem.Decode(rsaPublicKeyContent)
	if block == nil {
		return util.NULL, errors.New("pem.Decode：rsaPublicKey decode error")
	}
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509.ParsePKCS1PublicKey：%w", err)
	}
	cipherByte, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pubKey, []byte(text), nil)
	if err != nil {
		return "", fmt.Errorf("rsa.EncryptOAEP：%w", err)
	}
	return base64.StdEncoding.EncodeToString(cipherByte), nil
}

// 敏感信息解密，默认 PKCS1
func V3DecryptText(cipherText string, rsaPrivateKeyContent []byte) (text string, err error) {
	cipherByte, _ := base64.StdEncoding.DecodeString(cipherText)
	block, _ := pem.Decode(rsaPrivateKeyContent)
	if block == nil {
		return util.NULL, errors.New("pem.Decode：rsaPrivateKey decode error")
	}
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509.ParsePKCS1PrivateKey：%w", err)
	}
	textByte, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, cipherByte, nil)
	if err != nil {
		return "", fmt.Errorf("rsa.DecryptOAEP：%w", err)
	}
	return string(textByte), nil
}

// 解密普通支付回调中的加密订单信息
func V3DecryptNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}

// 解密普通退款回调中的加密订单信息
func V3DecryptRefundNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptRefundResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptRefundResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}

// 解密合单支付回调中的加密订单信息
func V3DecryptCombineNotifyCipherText(ciphertext, nonce, additional, apiV3Key string) (result *V3DecryptCombineResult, err error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(ciphertext)
	decrypt, err := aes.GCMDecrypt(cipherBytes, []byte(nonce), []byte(additional), []byte(apiV3Key))
	if err != nil {
		return nil, fmt.Errorf("aes.GCMDecrypt, err:%+v", err)
	}
	result = &V3DecryptCombineResult{}
	if err = json.Unmarshal(decrypt, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s), err:%+v", string(decrypt), err)
	}
	return result, nil
}
