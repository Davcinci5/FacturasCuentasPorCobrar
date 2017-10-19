package Crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"../../Modulos/Estructuras"
)

//GeneraHmacSHA256 regresa un slice de bytes
//correspondientes a la Mac generada por un mensaje y una llave.
func GeneraHmacSHA256(message, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	mac := h.Sum(nil)
	return mac
}

//CheckMacSha256 algoritmo Hmac de verificación de Mac.
func CheckMacSha256(Doc MoEstructuras.Send, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(Doc.X1X2X3)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(Doc.Y1Y2Y3, expectedMAC)
}

//EncryptAes64 método de encriptación Aes base64
func EncryptAes64(key, message []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(message)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

//DecryptAes64 método de desencriptación Aes base65
func DecryptAes64(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if len(text) < aes.BlockSize {
		fmt.Println(err)
		return nil, errors.New("documento demasiado pequeño")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func GetKey() string {
	return "12345678900000000000000000000000"
}

func CodificaSend(Document interface{}) (MoEstructuras.Send, error) {
	var Send MoEstructuras.Send
	llavePrivada := []byte(GetKey())

	Data, err := json.Marshal(Document)
	if err != nil {
		fmt.Println(err)
		return Send, err
	}

	CriptoMensaje, err := EncryptAes64(llavePrivada, Data)
	if err != nil {
		fmt.Println(err)
		return Send, err
	}

	Mac := GeneraHmacSHA256(CriptoMensaje, llavePrivada)

	Send = MoEstructuras.Send{CriptoMensaje, Mac}

	return Send, nil
}

func DecodificaSend(Doc MoEstructuras.Send) ([]byte, error) {
	llavePrivada := []byte(GetKey())

	if !CheckMacSha256(Doc, llavePrivada) {
		return nil, errors.New("No coincidieron las Mac")
	}

	DataBytes, err := DecryptAes64(llavePrivada, Doc.X1X2X3)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return DataBytes, nil
}
