/**
 * Created by angelina-zf on 17/2/27.
 */

// yeeCrypto
// 用于hash加密的包
// 依赖： "golang.org/x/crypto/bcrypt"
package yeeCrypto

import (
	"golang.org/x/crypto/bcrypt"
	"crypto/sha256"
	"encoding/hex"
	"crypto/sha512"
	"crypto/md5"
)

// Sha256Hex
func Sha256Hex(data []byte) string {
	out := sha256.Sum256(data)
	return hex.EncodeToString(out[:])
}

// Sha512Hex
func Sha512Hex(data []byte) string {
	out := sha512.Sum512(data)
	return hex.EncodeToString(out[:])
}

// Md5Hex 小写hex
func Md5Hex(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

// HashPassword
// 密码加密
func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, 0)
}

// CheckPasswordHash
// 加密后的密码的校验
func CheckPasswordHash(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
