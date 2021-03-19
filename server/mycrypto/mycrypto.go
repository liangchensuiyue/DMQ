package mycrypto

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

func GenerateRsaKey(keySize int, route string) {
	// 私钥生成
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	file, _ := os.Create(route + "/pri.pem")
	defer file.Close()
	writer := bufio.NewWriter(file)
	if err != nil {
		panic(err)
	}
	derPri := x509.MarshalPKCS1PrivateKey(privateKey)
	// fmt.Println(len(derPri), privateKey.Size(),privateKey.PublicKey.Size())
	block := pem.Block{
		Type:  "gds rsa private key abc", // 这个地方写字符串就行
		Bytes: derPri,
	}
	pem.Encode(writer, &block)
	writer.Flush()

	// 公钥生成
	derTextPub, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	// fmt.Println(len(derTextPub), privateKey.PublicKey.Size())
	block1 := pem.Block{
		Type:  "rsa public key",
		Bytes: derTextPub,
	}

	file1, _ := os.Create(route + "/pub.pem")
	defer file1.Close()
	writer1 := bufio.NewWriter(file1)
	pem.Encode(writer1, &block1)
	writer1.Flush()
}

func Encrypt(filepath string, msg string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	stat, _ := file.Stat()
	buf := make([]byte, stat.Size())
	file.Read(buf)
	block, _ := pem.Decode(buf)
	pubkeyinterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pubkey := pubkeyinterface.(*rsa.PublicKey)
	text, _ := rsa.EncryptPKCS1v15(rand.Reader, pubkey, []byte(msg))
	return base64.StdEncoding.EncodeToString(text), nil
}
func Decrypt(filepath string, msg string) (string, error) {
	file, _ := os.Open(filepath)
	stat, _ := file.Stat()
	buf := make([]byte, stat.Size())
	file.Read(buf)
	block, _ := pem.Decode(buf)
	prikeyinterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	prikey := prikeyinterface
	data, err1 := base64.StdEncoding.DecodeString(msg)
	if err1 != nil {
		return "", err1
	}
	text, err2 := rsa.DecryptPKCS1v15(rand.Reader, prikey, data)
	return string(text), err2
}
