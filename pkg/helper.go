package whatever

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	heritage "github.com/jmulhern/heritage/pkg"
	"gopkg.in/yaml.v3"
)

func OpenPacket() (packet heritage.Packet) {
	if raw, err := os.ReadFile("garden.yaml"); err == nil {
		if err := yaml.Unmarshal(raw, &packet); err == nil {
			fmt.Println("using garden.yaml")
			return packet
		}
	}
	panic("no packet loaded")
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
