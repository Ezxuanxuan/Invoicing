package cookie

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

const key = ":djasgDAGHH_ji1283"

//加密
func EncryptionId(id int) string {
	return strconv.Itoa(id) + key
}

//解密
func DecryptId(idValue string) int {
	ss := strings.Split(idValue, ":")
	id, err := strconv.Atoi(ss[0])
	if err != nil {
		return -1
	}
	return id
}

//MD5加密
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
