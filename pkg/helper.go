package whatever

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	what "github.com/jmulhern/what/pkg"
	"gopkg.in/yaml.v3"
)

func GetWhat() (what what.What) {
	if raw, err := os.ReadFile("what.yaml"); err == nil {
		if err := yaml.Unmarshal(raw, &what); err == nil {
			fmt.Println("using what.yaml")
			return what
		}
	}
	panic("no props loaded")
}

func PeekAt(a any) {
	raw, _ := json.MarshalIndent(a, "", " ")
	fmt.Println(string(raw))
}

func Encrypt(plaintext string) string {
	iv := []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
	block, err := aes.NewCipher([]byte(os.Getenv("KEY")))
	if err != nil {
		panic(err)
	}
	plainText := []byte(plaintext)
	encryptor := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	encryptor.XORKeyStream(cipherText, plainText)
	return base64.StdEncoding.EncodeToString(cipherText)
}

func Decrypt(cipherStr string) string {
	iv := []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
	block, err := aes.NewCipher([]byte(os.Getenv("KEY")))
	if err != nil {
		panic(err)
	}
	cipherText, err := base64.StdEncoding.DecodeString(cipherStr)
	if err != nil {
		panic(err)
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	decrypter.XORKeyStream(plainText, cipherText)
	return string(plainText)
}
