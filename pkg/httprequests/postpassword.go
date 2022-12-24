package httprequests

import (
	"crypto/aes"
	"crypto/cipher"
	cr "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var Password []byte
var LoaderCurrent LoaderKey

type LoaderKey struct {
	Cipher     string `json:"cipher"`
	Identifier string `json:"identifier"`
	Key        string `json:"key"`
	Timestamp  int64  `json:"timestamp"`
}

func (l *LoaderKey) ChangeKey() {
	key := []byte(RandStringBytesMaskImprSrcSB(32))
	plaintext := []byte(RandStringBytesMaskImprSrcSB(16))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(cr.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	timeNow, _ := strconv.Atoi(time.Now().Format("20060602150405"))
	l.Cipher = "aes-gcm"
	l.Key = string(ciphertext)
	l.Timestamp = int64(timeNow)
	l.Identifier = "wskey"
}

// GetNewPassword returns a new password to send data to post requests on websockets.
func GetNewPassword() []byte {
	return Password
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandStringBytesMaskImprSrcSB is used to get loader key.
//
// Thanks so much to best answer in Stack Overflow existence.
//
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}

// RefreshPassword refreshs password once in a time, to avoid anyone else except the code to access the websocket server.
//
// Password and LoaderKey is stroed in memory, and is refreshed every 5 minutes.
/* func RefreshPassword() {
	EncryptionKey := LoaderKey{}
	item := configstore.NewItem("encryption-key", `{"identifier":"storage","cipher":"aes-gcm","timestamp":1559309532,"key":"b6a942c0c0c75cc87f37d9e880c440ac124e040f263611d9d236b8ed92e35521"}`, 0)
	fmt.Println(LoaderKey)
	k, err := keyloader.LoadKey(LoaderKey)
	Password, err = k.Encrypt([]byte(RandStringBytesMaskImprSrcSB(16)))
	if err != nil {
		panic(err)
	}
	var refreshDelay = 5 * time.Minute
	for {
		LoaderKey = RandStringBytesMaskImprSrcSB(32)
		k, err = keyloader.LoadKey(LoaderKey)
		if err != nil {
			panic(err)
		}
		Password, err = k.Encrypt(Password)
		if err != nil {
			panic(err)
		}
		log.Println("Password and loading key refreshed.")
		time.Sleep(refreshDelay)
	}
} */

func RefreshPassword() {
	LoaderCurrent.ChangeKey()
	fmt.Println(LoaderCurrent)
}
