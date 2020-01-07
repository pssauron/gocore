//============================================================
// 描述: 字符串处理
// 作者: Simon
// 日期: 2019/10/29 3:51 下午
//
//============================================================

package strutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

//EncodePassword 创建Password
func EncodePassword(password string) string {
	en, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return base64.StdEncoding.EncodeToString(en)
}

//ComparePassword 密码比对
func ComparePassword(password, encodePwd string) bool {
	bb, err := base64.StdEncoding.DecodeString(encodePwd)

	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword(bb, []byte(password))

	if err != nil {
		return false
	}

	return true

}

//Index 支持UTF8 index
func Index(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

func SubStr(str string, begin, length int) string {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

//Match 正则匹配
func Match(str string, exp string) bool {

	reg := regexp.MustCompile(exp)

	return reg.Match([]byte(str))

}

func ToSnakeCase(str string) string {

	var snake string

	for k, v := range str {
		if k == 0 {
			snake = strings.ToLower(string(str[0]))
		} else {
			if unicode.IsUpper(rune(v)) {
				snake += "_" + strings.ToLower(string(v))
			} else {
				snake += strings.ToLower(string(v))
			}

		}

	}
	return snake
}

func ToCamelCase(str string) string {
	var camel string
	var toUpper bool
	str = strings.TrimLeft(str, "_")

	for k, v := range str {
		if k == 0 {
			camel = strings.ToUpper(string(v))
		} else {
			if v == '_' {
				toUpper = true
			} else {
				if toUpper {
					camel += strings.ToUpper(string(v))
					toUpper = false
				} else {
					camel += string(v)
				}
			}
		}
	}
	return camel
}

// ============================== AES =======================================
// =========================AES 秘钥长度必须是 16 或 16 的倍数 ===================
func PKSC7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origDatastr, keystr string) (string, error) {

	for len(keystr)%16 != 0 {
		keystr += "0"
	}

	origData := []byte(origDatastr)
	key := []byte(keystr)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKSC7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func AesDecrypt(cryptedstr, keystr string) (string, error) {
	for len(keystr)%16 != 0 {
		keystr += "0"
	}
	crypted, err := base64.StdEncoding.DecodeString(cryptedstr)
	if err != nil {
		return "", errors.New("Base64 error")
	}
	key := []byte(keystr)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return string(origData), nil
}

// ============================== AES END ===================================

// ============================== MD5 Start =================================

func MD5(crypted string) string {
	h := md5.New()
	h.Write([]byte(crypted))
	return hex.EncodeToString(h.Sum(nil))
}

// ================================ MD5 END =================================
