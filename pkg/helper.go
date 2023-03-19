package whatever

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	what "github.com/jmulhern/what/pkg"
	"gopkg.in/yaml.v3"
)

func EncryptWhat() {
	whatRaw, err := os.ReadFile("what.yaml")
	if err != nil {
		panic(err)
	}
	keyRaw, err := os.ReadFile("what.key")
	if err != nil {
		panic(err)
	}
	encrypted, err := Encrypt(whatRaw, keyRaw)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("what.encrypted", encrypted, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func GetWhat() (what what.What) {
	if raw, err := os.ReadFile("what.yaml"); err == nil {
		if err := yaml.Unmarshal(raw, &what); err == nil {
			fmt.Println("using what.yaml")
			return what
		}
	}
	if raw, err := os.ReadFile("what.encrypted"); err == nil {
		if plain, err := Decrypt(raw, GetKey()); err == nil {
			if err := yaml.Unmarshal(plain, &what); err == nil {
				fmt.Println("using what.encrypted")
				return what
			}
		}
	}
	panic("no props loaded")
}

func GetKey() string {
	if key := os.Getenv("KEY"); key != "" {
		return key
	}
	return ""
}

func PeekAt(a any) {
	raw, _ := json.MarshalIndent(a, "", " ")
	fmt.Println(string(raw))
}

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	cipherText := make([]byte, aes.BlockSize+len(plaintext))

	// Slice of first 16 bytes
	iv := cipherText[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	return cipherText, nil
}

func Decrypt(cipherText []byte, key string) ([]byte, error) {
	// Create the AES cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("test too short")
	}

	// Get the 16 byte IV
	iv := cipherText[:aes.BlockSize]

	// Remove the IV from the ciphertext
	cipherText = cipherText[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}
