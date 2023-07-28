package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateCt0() string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100000)
	randomStr := strconv.Itoa(randomNum)
	hash := md5.Sum([]byte(randomStr))
	return hex.EncodeToString(hash[:])
}
func GetHeaders(authToken string) map[string]string {
	headers := map[string]string{
		"authorization":             "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
		"referer":                   "https://twitter.com/",
		"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		"x-twitter-auth-type":       "",
		"x-twitter-active-user":     "yes",
		"x-twitter-client-language": "en",
	}
	ct0 := GenerateCt0()
	headers["x-csrf-token"] = ct0
	headers["cookie"] = fmt.Sprintf("auth_token=%s; ct0=%s", authToken, ct0)
	headers["x-twitter-auth-type"] = "OAuth2Session"

	return headers
}
