package main

import (
		"os"
		"fmt"
		"crypto/aes"
     		"crypto/cipher"
    		"crypto/rand"
    		"encoding/hex"
		"io"
)

func main() {

	cbc_encrypt()
	cbc_decrypt()

}

func cbc_encrypt() {

		key := []byte(os.Getenv("GOKEY"))

		// Block         1         2         3         4
		// 	1234567890123456789012345678901234567890
		//	this is our text
		// Plain text MUST be a full block!
		plaintext := []byte("this is our text")
		fmt.Printf("Key : %s\n", key)

		block, err := aes.NewCipher(key)
		if err != nil {
			panic(err.Error())
		}

		ciphertext := make([]byte, aes.BlockSize+len(plaintext))
   		iv := ciphertext[:aes.BlockSize]
   		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
   			panic(err)
   		}

   		mode := cipher.NewCBCEncrypter(block, iv)
   		mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

   		fmt.Printf("Encrypt : %x\n", ciphertext)
}

func cbc_decrypt() {

		key := []byte(os.Getenv("GOKEY"))

		fmt.Printf("Key : %s\n", key)
		ciphertext, _ := hex.DecodeString("a3f0015bc11fd02e08268540260857a495b64652c6c21c8447736acd960a6709")

    		block, err := aes.NewCipher(key)
    		if err != nil {
    			panic(err.Error())
    		}

    		if len(ciphertext) < aes.BlockSize {
    			panic("ciphertext too short")
    		}
    		iv := ciphertext[:aes.BlockSize]
    		ciphertext = ciphertext[aes.BlockSize:]

    		if len(ciphertext)%aes.BlockSize != 0 {
    			panic("ciphertext is not a multiple of the block size")
    		}

    		mode := cipher.NewCBCDecrypter(block, iv)

    		mode.CryptBlocks(ciphertext, ciphertext)

   		fmt.Printf("Decrypt : %s\n", ciphertext)
}
