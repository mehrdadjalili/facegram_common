package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
)

func (e *Encryption) RsaSign(data []byte) (string, error) {

	key, err := getPrivateKey(e.privateKey)

	if err != nil {
		return "", derrors.InternalError()
	}

	msgHash := sha256.New()
	_, err = msgHash.Write(data)
	if err != nil {
		return "", derrors.InternalError()
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, key, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		return "", derrors.InternalError()
	}

	return base64.URLEncoding.EncodeToString(signature), nil
}

func (e *Encryption) RsaVerify(signature string, data []byte) (bool, error) {

	ciphertext, _ := base64.URLEncoding.DecodeString(signature)
	key, err := getPublicKey(e.publicKey)
	if err != nil {
		return false, derrors.InternalError()
	}

	msgHash := sha256.New()
	_, err = msgHash.Write(data)
	if err != nil {
		return false, derrors.InternalError()
	}

	msgHashSum := msgHash.Sum(nil)

	err = rsa.VerifyPSS(key, crypto.SHA256, msgHashSum, ciphertext, nil)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (e *Encryption) RsaEncrypt(plainText string) (string, error) {
	bytes := []byte(e.publicKey)
	publicKey, err := convertBytesToPublicKey(bytes)
	if err != nil {
		return "", derrors.InternalError()
	}
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	if err != nil {
		return "", derrors.InternalError()
	}
	return base64.URLEncoding.EncodeToString(cipher), nil
}

func (e *Encryption) RsaDecrypt(encryptedMessage string) (string, error) {
	bytes := []byte(e.privateKey)
	privateKey, err := convertBytesToPrivateKey(bytes)
	if err != nil {
		return "", derrors.InternalError()
	}

	msg, err := base64.URLEncoding.DecodeString(encryptedMessage)
	if e != nil {
		return "", derrors.InternalError()
	}

	plainMessage, e5 := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		msg,
	)
	if e5 != nil {
		return "", e5
	}
	return string(plainMessage), nil
}

func getPrivateKey(key string) (*rsa.PrivateKey, error) {
	privateKey, err := convertBytesToPrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func getPublicKey(key string) (*rsa.PublicKey, error) {
	pubKey, err := convertBytesToPublicKey([]byte(key))
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}

func loadKey(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func convertBytesToPrivateKey(keyBytes []byte) (*rsa.PrivateKey, error) {
	var err error

	block, _ := pem.Decode(keyBytes)
	blockBytes := block.Bytes
	ok := x509.IsEncryptedPEMBlock(block)

	if ok {
		blockBytes, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(blockBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func convertBytesToPublicKey(keyBytes []byte) (*rsa.PublicKey, error) {
	var err error

	block, _ := pem.Decode(keyBytes)
	blockBytes := block.Bytes
	ok := x509.IsEncryptedPEMBlock(block)

	if ok {
		blockBytes, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}

	pkey, err := x509.ParsePKIXPublicKey(blockBytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := pkey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("got unexpected key type")
	}

	return rsaKey, nil
}
