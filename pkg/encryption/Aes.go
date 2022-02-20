package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
)

func (e *Encryption) AesEncrypt(data string) (string, error) {

	bKey, er := hex.DecodeString(e.aesKey)
	if er != nil {
		return "", derrors.InternalError()
	}
	plaintext := []byte(data)

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", derrors.InternalError()
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", derrors.InternalError()
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func (e *Encryption) AesDecrypt(data string) (string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(data)
	bKey, er := hex.DecodeString(e.aesKey)
	if er != nil {
		return "", derrors.InternalError()
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		return "", derrors.InternalError()
	}

	if len(ciphertext) < aes.BlockSize {
		return "", derrors.InternalError()
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}
