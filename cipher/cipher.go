package cipher

import (
    "encoding/gob"
	"encoding/base64"
	"encoding/hex"
    "encoding/json"
    "crypto/md5"
    "crypto/sha256"
	"crypto/sha512"
    "crypto/aes"
	"crypto/des"
	"crypto/cipher"
	"crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "crypto"
    //"errors"
    "bytes"
	//"fmt"
)

//-------------------------------------------- encode --------------------------------------------
func GobEncode(data interface{}) ([]byte, error) {
    buf := bytes.NewBuffer(nil)
    enc := gob.NewEncoder(buf)
    err := enc.Encode(data)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), err
}

func GobDecode(data []byte, to interface{}) error {
    buf := bytes.NewBuffer(data)
    dec := gob.NewDecoder(buf)
    return dec.Decode(to)
}

func Base64Encode(plainText []byte) string {
	return base64.StdEncoding.EncodeToString(plainText)
}

func Base64Decode(str string) []byte {
	data,err := base64.StdEncoding.DecodeString(str)
	if err!=nil {
		return nil
	}
	return data
}

func HexEncode(s []byte) string {
	return hex.EncodeToString(s)
}

func HexDecode(s string) []byte {
	data,err := hex.DecodeString(s)
	if err != nil {
		return nil
	}
	return data
}

func JsonEncode(v interface{}) []byte {
    j,err := json.Marshal(v)
    if err != nil {
        return nil
    }
    return j
}

func JsonDecode(data []byte,v interface{}) error {
    return json.Unmarshal(data,&v)
}
//---------------------------------------- Hash -------------------------------------------------------

func Sha256(plaintext []byte) []byte {
    cipher := sha256.Sum256(plaintext)
    return cipher[:]
}

func Sha512(plainText []byte) []byte {
    cipher:=sha512.Sum512(plainText)
	return cipher[:]
}

func Md5(data []byte) []byte {
    cipher := md5.Sum(data)
    return cipher[:]
}

func GenerateRsaKey(bits int) *rsa.PrivateKey {
    priv,err := rsa.GenerateKey(rand.Reader,bits)
    if err != nil {
        return nil
    }
    return priv
}

func MarshalRsaKey(priv *rsa.PrivateKey) (derPublic,derPrivate []byte,err error) {
    derPrivate = x509.MarshalPKCS1PrivateKey(priv)
    pub :=&priv.PublicKey
    err = nil
    derPublic,err = x509.MarshalPKIXPublicKey(pub)
    if err != nil {
        return nil,nil,err
    }
    return
}

func ParseRsaPriv(mpriv []byte) (*rsa.PrivateKey) {
    priv, err := x509.ParsePKCS1PrivateKey(mpriv)
    if err != nil {
        return nil
    }
    return priv
}

func ParseRsaPub(mpub []byte) (*rsa.PublicKey) {
    pubInterface, err := x509.ParsePKIXPublicKey(mpub)
    if err != nil {
        return nil
    }
    pub := pubInterface.(*rsa.PublicKey)
    return pub
}

func RsaEncrypt(pub *rsa.PublicKey,plaintext []byte) ([]byte, error) {
    return rsa.EncryptPKCS1v15(rand.Reader, pub, plaintext)
}

func RsaDecrypt(priv *rsa.PrivateKey,ciphertext []byte) ([]byte, error) {
    return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func Signature(priv *rsa.PrivateKey,msg []byte) ([]byte,error) {
    hashed := Sha256(msg)
    return rsa.SignPKCS1v15(rand.Reader,priv,crypto.SHA256,hashed)
}

func VerifySign(pub *rsa.PublicKey,sig []byte,msg []byte) error {
    hashed := Sha256(msg)
    return rsa.VerifyPKCS1v15(pub,crypto.SHA256,hashed,sig)
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{0}, padding)
    return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
    return bytes.TrimFunc(origData,
        func(r rune) bool {
            return r == rune(0)
        })
}

// Des Encryption Algorithm, key required (length:8)
func DesEncrypt(origData, key []byte) ([]byte, error) {
     block, err := des.NewCipher(key)
     if err != nil {
          return nil, err
     }
     origData = ZeroPadding(origData, block.BlockSize())
     blockMode := cipher.NewCBCEncrypter(block, key)
     crypted := make([]byte, len(origData))
     blockMode.CryptBlocks(crypted, origData)
     return crypted, nil
}
// Des Decryption Algorithm, key required (length:8)
func DesDecrypt(crypted, key []byte) ([]byte, error) {
     block, err := des.NewCipher(key)
     if err != nil {
          return nil, err
     }
     blockMode := cipher.NewCBCDecrypter(block, key)
     origData := make([]byte, len(crypted))
     blockMode.CryptBlocks(origData, crypted)
     origData = ZeroUnPadding(origData)
     return origData, nil
}

// 3Des Encryption Algorithm, key required (length:24)
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
     block, err := des.NewTripleDESCipher(key)
     if err != nil {
          return nil, err
     }
     origData = ZeroPadding(origData, block.BlockSize())
     blockMode := cipher.NewCBCEncrypter(block, key[:8])
     crypted := make([]byte, len(origData))
     blockMode.CryptBlocks(crypted, origData)
     return crypted, nil
}
// 3Des Encryption Algorithm, key required (length:24)
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
     block, err := des.NewTripleDESCipher(key)
     if err != nil {
          return nil, err
     }
     blockMode := cipher.NewCBCDecrypter(block, key[:8])
     origData := make([]byte, len(crypted))
     blockMode.CryptBlocks(origData, crypted)
     origData = ZeroUnPadding(origData)
     return origData, nil
}

func AesEncrypt(plantText, key []byte) ([]byte, error) {
   block, err := aes.NewCipher(key)
   if err != nil {
      return nil, err
   }
   plantText = ZeroPadding(plantText, block.BlockSize())

   blockMode := cipher.NewCBCEncrypter(block, key)

   ciphertext := make([]byte, len(plantText))

   blockMode.CryptBlocks(ciphertext, plantText)
   return ciphertext, nil
}

func AesDecrypt(ciphertext, key []byte) ([]byte, error) {
   block, err := aes.NewCipher(key)
   if err != nil {
      return nil, err
   }
   blockModel := cipher.NewCBCDecrypter(block, key)
   plantText := make([]byte, len(ciphertext))
   blockModel.CryptBlocks(plantText, ciphertext)
   plantText = ZeroUnPadding(plantText)
   return plantText, nil
}