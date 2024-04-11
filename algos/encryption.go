package algos

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"

	// "encoding/base64"
	"fmt"
	"io"
)

// Encrypt encrypts a plaintext value using AES encryption with a given key.
func Encrypt(plaintext []byte, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	encodedCiphertext := Encode(ciphertext)
	// decodedCiphertext := Decode(encodedCiphertext)
	// fmt.Println(string(ciphertext))
	// // fmt.Println(string(decodedCiphertext))
	// if len(decodedCiphertext) < aes.BlockSize {
	// 	return "", fmt.Errorf("decodedCiphertext too short")
	// }

	// iv = decodedCiphertext[:aes.BlockSize]
	// decodedCiphertext = decodedCiphertext[aes.BlockSize:]

	// stream = cipher.NewCFBDecrypter(block, iv)
	// stream.XORKeyStream(decodedCiphertext, decodedCiphertext)

	return encodedCiphertext, nil
}

// Decrypt decrypts a ciphertext value using AES encryption with a given key.
func Decrypt(ciphertext string, key string) (string, error) {
	decodedCiphertext := Decode(ciphertext)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(decodedCiphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := decodedCiphertext[:aes.BlockSize]
	decodedCiphertext = decodedCiphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decodedCiphertext, decodedCiphertext)

	return string(decodedCiphertext), nil
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(d string) []byte {
	decodedString, _ := base64.StdEncoding.DecodeString(d)
	return decodedString
}
