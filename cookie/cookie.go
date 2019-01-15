package cookie

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"strconv"
)

const (
	base64Table = "0123456789+abcdefghijklmnopqrstuvwxyz/ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var coder = base64.NewEncoding(base64Table)

//加密
func EncryptionId(id int64) string {
	temp := strconv.FormatInt(id, 10)
	src := []byte(temp)
	return string(coder.EncodeToString(src))
}

//解密
func DecryptId(idValue string) int64 {
	re, err := coder.DecodeString(idValue)
	if err != nil {
		return 0
	}
	temp, err := strconv.ParseInt(string(re), 10, 64)
	if err != nil {
		return 0
	}
	return temp
}

//MD5加密
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
