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

const privateKeyPEM = `
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQClVt/r2PokXInK
4ftRJdo+uxzRfatXOgS6cFhKoF3YQEp7R/kI01p2YtgbG7vLFWKV4dl2+tK+ROsN
SzEiB/9Zpz0FkqhYeRLWayOV5d9Vgcxbg/KhfLLyfuVRBm1heK6HxlIRLUt9OAp/
Ap1U6nPK/nsn2aVUxHRiKt4I+VKZMMSeR5s0re6Y7MzAr3iW5tFd+dE1ky2qM7pr
ur+B6woa6shi4J5GG/72tJ42Ppr9tCduO1IaWFah/1zaWEOIL6+5smKasRS06hJz
Qd0ZL6p5w/48OhM2+lDnQEiy9TdOZLd+pLnC12+cTOYT4L7LCbmpVkdaUbqaEB+6
jOj0EG77AgMBAAECggEAALpND/kHqAeEAKDnuhkXGBZuT2vnSENDsTkvufrsHtJU
qUZfH7Me7sCBJTP1076gc7gTJb5rKbnpaRqg9kYqDrQahSfRwf3u6dU2c3cIkxQd
UVW4uTbMt2XEWSVpHbgiZ5Ddrl30K86gCilvXT7wS1YjiDFPoUiV68FA/OUBAF4g
Q3fbmpDeOKpYvs7wxpf5jhT6nmZb46/UWRHo+JQgoNjTGKqS3S4PdJ9phxgOac9F
lathck4lDfCugduo8Or0G5UjruRgp+6s1fH30ie9md3hAH01CpVbe+eIWprGQCDc
yXfoQ8lZKa54sVM/9H2wIi5rYFoxVFTcV3KdIEdRgQKBgQDjuERy1tyf8jNXYCC7
v/BpIN1g1EyACnhnEf9hujtSwuQYk5gX00hr3XgrMfSdP3VTb7MkvHIoj+Rdw8ie
eZMTdqwpVQ+T6tLAXRNWBBOopX3m+pyPa98fZQKNzl/P25QS1pUevcfpkEVcXwLz
y3XXV0D4ouPvvUkJ8LymldgIgQKBgQC532P86Z6OIQdVXJOUYExXnEDcNXRtlDIS
66tkF1RWAsy21CxZBANOjCyttWgxudbJOB/HVnzWXc9prHdmWbJUrotEsUNDokTt
9jq73ch2wYM/aomSrPyRrG88E2Gmr+67iyuDgdzt8YRSUEKpqgpbhe+A1JnTxNer
Ud3zFo/ZewKBgBwFEGcRCRSlOKwHp/9yHqLQ6GyBsugOYcJM9J+RyrrkQCzF/HDy
Dnc3SRIHk5HFvSoHFIPwrBtRmUfwTz8wtmgusgBj9wa8XjJNQZPT7JdTxaomLB15
qHq0cxv/yMpKum2W+cJOl8qldeNnzXUyE83rbpMpd+KH5/TRKmVVCsABAoGAEpF5
eKOR/lrYU7O63oC8P6hRZm5EoknCstcuOQKHn0wKTV1mzMG03tzr/bJ4pTcOeO2N
ZymBsRyQAtuC1gux3/nL2eHneVM1lZwag/gE9bAhP22SLr/vP1I9jn/VKoS60at2
fl2zx2VwNZTlA/QDst7vbSxP7bLlZKz6AjXHyw8CgYEAzLyY2HRMGwvy8aS3qf+j
d0Uzb8tV7ZUZ81hfdg1u1Ig8b+2hCRKzVcPqQ5osUZJ7oGhPxDSv7Y2wGCFdltHl
irvva3gSfkQsY8Kq7cjlD4bdo43LmXOKRp3E5RzV3defyzTFZ5amI45YP/d3bSPY
lSMayVa5NfMLrwKeYrTOYHI=
-----END PRIVATE KEY-----
`

func SignFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(data)

	keyData := []byte(privateKeyPEM)

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

const publicKeyPEM = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApVbf69j6JFyJyuH7USXa
Prsc0X2rVzoEunBYSqBd2EBKe0f5CNNadmLYGxu7yxVileHZdvrSvkTrDUsxIgf/
Wac9BZKoWHkS1msjleXfVYHMW4PyoXyy8n7lUQZtYXiuh8ZSES1LfTgKfwKdVOpz
yv57J9mlVMR0YireCPlSmTDEnkebNK3umOzMwK94lubRXfnRNZMtqjO6a7q/gesK
GurIYuCeRhv+9rSeNj6a/bQnbjtSGlhWof9c2lhDiC+vubJimrEUtOoSc0HdGS+q
ecP+PDoTNvpQ50BIsvU3TmS3fqS5wtdvnEzmE+C+ywm5qVZHWlG6mhAfuozo9BBu
+wIDAQAB
-----END PUBLIC KEY-----
`

func GetPublicKeyPEM() (string, error) {
	return publicKeyPEM, nil
}
