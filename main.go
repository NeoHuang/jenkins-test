package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {

	text := []byte(`{
"provider": "digital_turbine",
"label": "some_random_label",
"tracker": "abc123",
"campaign": "test_campaign",
"adgroup": "test adgroup",
"creative": "test creative- %3 user",
"install_callback": "http://my.server.com/postback?source=adjust&os_ver={os_version}&app={app_id}&click_id={click_id}",
"created_at":1579100361
}`)
	key := []byte("AES256Key-32Characters1234567890")

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	//	simple(c, text)
	gcm(c, text)
	//	cfb(c, text)

}

func simple(c cipher.Block, text []byte) {
	fmt.Println("Simple encode")
	out := make([]byte, len(text))
	c.Encrypt(out, text)
	decoded := make([]byte, len(text))
	fmt.Println("encode", base64.URLEncoding.EncodeToString(out))
	c.Decrypt(decoded, out)
	fmt.Println("decode", string(decoded))
}

func gcm(c cipher.Block, text []byte) {
	fmt.Println("GCM")
	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce, err := hex.DecodeString("b4e3a92c421532044a2d2e97")
	// populates our nonce with a cryptographically secure
	// random sequence
	// if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
	// 	fmt.Println(err)
	// }

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	sealed := gcm.Seal(nonce, nonce, text, nil)
	encoded := base64.URLEncoding.EncodeToString(sealed)
	fmt.Println("encode", encoded)

	gcm, err = cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := sealed[:nonceSize], sealed[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("decode", string(plaintext))
}

func cfb(c cipher.Block, text []byte) {
	fmt.Println("CFB")
	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(c, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], text)

	//returns to base64 encoded string
	fmt.Println(base64.URLEncoding.EncodeToString(cipherText))
}
