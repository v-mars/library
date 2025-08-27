package aes

import "testing"

const (
	key = "IgkibX71IEf38212"
	iv  = "IgkibX71IEf38212"
)

func TestEncrypt(t *testing.T) {
	t.Log(New(key, iv).Encrypt("123456"))
}

func TestDecrypt(t *testing.T) {
	t.Log(New(key, iv).Decrypt("GO-ri84zevE-z1biJwfQPw=="))
}

func BenchmarkEncryptAndDecrypt(b *testing.B) {
	b.ResetTimer()
	aes := New(key, iv)
	for i := 0; i < b.N; i++ {
		encryptString, _ := aes.Encrypt("123456")
		aes.Decrypt(encryptString)
	}
}
