package object

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/labstack/gommon/log"
	"gopkg.in/square/go-jose.v2"
	"io"
	"math/big"
	"os"
	"simple-casdoor/util"
	"time"
)

type Cert struct {
	Certificate []byte `json:"certificate"`
	PrivateKey  []byte `json:"privateKey"`
}

var cert *Cert

func init() {
	cert = new(Cert)
	permPath, _ := util.GetAbsolutePath("token_jwt_key.pem")
	keyPath, _ := util.GetAbsolutePath("token_jwt_key.key")

	pemData, err := getFileData(permPath)
	if err != nil {
		// 生成 cert
		certStr, keyStr := generateRsaKeys(4096, 20, "fbCert", "fireboom")
		cert.Certificate = []byte(certStr)
		cert.PrivateKey = []byte(keyStr)
		// 生成文件
		go generateCertFile(permPath, cert.Certificate)
		go generateCertFile(keyPath, cert.PrivateKey)
		log.Info("已生成cert文件...")
	} else {
		cert.Certificate = pemData
		key, err := getFileData(keyPath)
		if err != nil {
			log.Error(err.Error())
		}
		cert.PrivateKey = key
	}
}

func GetJsonWebKeySet() (jose.JSONWebKeySet, error) {
	jwks := jose.JSONWebKeySet{}

	certPemBlock := cert.Certificate
	certDerBlock, _ := pem.Decode(certPemBlock)
	x509Cert, _ := x509.ParseCertificate(certDerBlock.Bytes)

	var jwk jose.JSONWebKey
	jwk.Key = x509Cert.PublicKey
	jwk.Certificates = []*x509.Certificate{x509Cert}
	jwk.KeyID = "fireboom"
	jwk.Algorithm = "RS256"
	jwk.Use = "sig"

	jwks.Keys = append(jwks.Keys, jwk)

	return jwks, nil
}

func getFileData(filePath string) ([]byte, error) {
	// 打开文本文件进行读取
	file, err := os.Open(filePath)
	if err != nil {
		log.Error("无法打开文件:%v", err)
		return nil, err
	}
	defer file.Close()

	// 读取文件的所有数据
	data, err := io.ReadAll(file)
	if err != nil {
		log.Error("读取文件时出错:%v", err)
		return nil, err
	}

	// 将数据转换为字节数组
	return data, nil
}

func generateRsaKeys(bitSize int, expireInYears int, commonName string, organization string) (string, string) {
	// Generate RSA key.
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}

	// Encode private key to PKCS#1 ASN.1 PEM.
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	tml := x509.Certificate{
		// you can add any attr that you need
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(expireInYears, 0, 0),
		// you have to generate a different serial number each execution
		SerialNumber: big.NewInt(123456),
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{organization},
		},
		BasicConstraintsValid: true,
	}
	cert, err := x509.CreateCertificate(rand.Reader, &tml, &tml, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}

	// Generate a pem block with the certificate
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})

	return string(certPem), string(privateKeyPem)
}

func generateCertFile(filename string, content []byte) {
	file, err := os.Create(filename)
	if err != nil {
		log.Error(err.Error())
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		log.Error(err.Error())
	}
}
