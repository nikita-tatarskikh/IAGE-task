package task2

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

type HMAC struct {

}

func (sec *HMAC) CreateSign(request *Request) string {
	var s, key string
	s = request.GetS()
	key = request.GetKey()
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(s))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}


func NewHMAC() *HMAC {
	return &HMAC{}
}

