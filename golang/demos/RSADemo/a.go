// Package main provides ...
package newDir

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func JustDemo() {
	println("//<<-------------------------JustDemo start-----------")
	start := time.Now()

	fmt.Printf("JustDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------JustDemo end----------->>")
}

func GenerateRSAKeyDemo() {
	println("//<<-------------------------GenerateRSAKeyDemo start-----------")
	start := time.Now()
	if err := GenRsaKey(1024); err != nil {
		log.Fatal("密钥文件生成失败！")
	}
	log.Println("密钥文件生成成功！")
	fmt.Printf("GenerateRSAKeyDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------GenerateRSAKeyDemo end----------->>")
}

// func main() {
// 	var bits int
// 	flag.IntVar(&bits, "b", 1024, "密钥长度，默认为1024位")
// 	if err := GenRsaKey(bits); err != nil {
// 		log.Fatal("密钥文件生成失败！")
// 	}
// 	log.Println("密钥文件生成成功！")
// }

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

const msg = "win27v@gmail.com"

func DecryptByPrivateKeyDemo() {
	println("//<<-------------------------DecryptByPrivateKeyDemo start-----------")
	start := time.Now()
	privateKey, err := ioutil.ReadFile("./private.pem")
	if err != nil {
		panic(err.Error())
	}
	block, result := pem.Decode(privateKey)
	if block == nil {
		panic("private key error!")
	}

	block2, _ := pem.Decode(result)
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic("private key error!")
	}
	// log.Printf("block2.Bytes: %+v\n", string(block2.Bytes))
	// log.Printf("block2: %+v\n", block2)
	// log.Printf("priv: %+v\n", priv)
	ciphertext, _ := EncryptByPublicKeyDemo()
	bs, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		panic(err.Error())
	}
	if string(bs) != msg {
		panic("data not match after encrypt")
	}
	log.Printf("string(bs): %+v\n", string(bs))
	fmt.Printf("DecryptByPrivateKeyDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------DecryptByPrivateKeyDemo end----------->>")
}

func EncryptByPublicKeyDemo() ([]byte, error) {
	println("//<<-------------------------EncryptByPublicKeyDemo start-----------")
	start := time.Now()
	publicKey, err := ioutil.ReadFile("./public.pem")
	if err != nil {
		panic(err.Error())
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		panic("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err.Error())
	}
	pub := pubInterface.(*rsa.PublicKey)
	fmt.Printf("EncryptByPublicKeyDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------EncryptByPublicKeyDemo end----------->>")
	return rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(msg))
}
