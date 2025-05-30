package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// SignFileHash - функция подписи договора
func SignFileHash(filePath, privateKeyPath string) ([]byte, error) {
	//nolint:gosec // trusted file path from config
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(data)

	//nolint:gosec // trusted file path from config
	keyData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("invalid PEM block or not a PRIVATE KEY")
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA private key")
	}

	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
}

// LoadPublicKeyPEM - функция, которая выгружает публичный ключ
func LoadPublicKeyPEM(path string) (string, error) {
	//nolint:gosec // trusted file path from config
	pubKey, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(pubKey), nil
}
