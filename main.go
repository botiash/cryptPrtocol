package main

import (
	"fmt"
)

// Функция для шифрования сообщения с использованием XOR
func xorEncrypt(message []byte, key []byte) []byte {
	encrypted := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		encrypted[i] = message[i] ^ key[i%len(key)]
	}
	return encrypted
}

// Функция для расшифрования сообщения с использованием XOR
func xorDecrypt(encrypted []byte, key []byte) []byte {
	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i++ {
		decrypted[i] = encrypted[i] ^ key[i%len(key)]
	}
	return decrypted
}

func main() {
	message := []byte("Hello, World!")
	key := []byte("SecretKey")

	encrypted := xorEncrypt(message, key)
	decrypted := xorDecrypt(encrypted, key)

	fmt.Println("Original Message:", string(message))
	fmt.Println("Encrypted Message:", string(encrypted))
	fmt.Println("Decrypted Message:", string(decrypted))
}
