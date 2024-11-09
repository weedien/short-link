package toolkit

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash/fnv"
	"math"
)

var CHARS = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}
var SIZE = len(CHARS)

// MD5 计算字符串的MD5哈希值
func MD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA1 计算字符串的SHA1哈希值
func SHA1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA256 计算字符串的SHA256哈希值
func SHA256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA512 计算字符串的SHA512哈希值
func SHA512(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func convertDecToBase62(num int64) string {
	if num == 0 {
		return string(CHARS[0])
	}
	var sb []rune
	for num > 0 {
		i := num % int64(SIZE)
		sb = append([]rune{CHARS[i]}, sb...)
		num /= int64(SIZE)
	}
	return string(sb)
}

func HashToBase62(str string) string {
	h := fnv.New32a()
	_, err := h.Write([]byte(str))
	if err != nil {
		return ""
	}
	i := int64(h.Sum32())
	if i < 0 {
		i = math.MaxInt32 - i
	}
	return convertDecToBase62(i)
}
