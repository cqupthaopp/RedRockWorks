package Utils

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}
	unPadNum := int(dst[len(dst)-1])
	return dst[:(len(dst) - unPadNum)], nil
}

const desKey = "cqupthao"

func DesEncoding(src string) string {

	srcByte := []byte(src)
	block, err := des.NewCipher([]byte(desKey))
	if err != nil {
		fmt.Println(err)
	}

	newSrcByte := PadPwd(srcByte, block.BlockSize())
	dst := make([]byte, len(newSrcByte))
	block.Encrypt(dst, newSrcByte)

	pwd := base64.StdEncoding.EncodeToString(dst)

	return pwd
}

func DesDecoding(pwd string) string {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd
	}
	block, errBlock := des.NewCipher([]byte(desKey))
	if errBlock != nil {
		return pwd
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, _ = UnPadPwd(dst)
	return string(dst)
}

func GetNewToken(s interface{}) string {

	data, err := json.Marshal(&s)

	if err != nil {
		fmt.Println(err)
	}

	code := DesEncoding(string(data))

	return code

}

func PraiseToken(s interface{}, code string) {

	json.Unmarshal([]byte(DesDecoding(code)), s)

}

func CreateNewToken(name string) string {

	token := Token{
		name, time.Now().Unix(), "null",
	}

	code := GetNewToken(token)

	return code

}
