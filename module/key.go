package module

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// 生成密钥
func generateKey() []byte {
	key := make([]byte, 32)
	rand.Read(key)
	return key
}

// 加密
func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用CBC模式，需要初始化向量
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 随机生成nonce
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	// 加密
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func MustEncrypt(plaintext []byte, key []byte) []byte {
	ciphertext, err := encrypt(plaintext, key)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// 解密
func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用CBC模式，需要初始化向量
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 获取nonce
	nonceSize := aesGCM.NonceSize()
	nonce := ciphertext[:nonceSize]
	ciphertext = ciphertext[nonceSize:]

	// 解密
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
func MustDecrypt(ciphertext []byte, key []byte) []byte {
	plaintext, err := decrypt(ciphertext, key)
	if err != nil {
		panic(err)
	}
	return plaintext
}
