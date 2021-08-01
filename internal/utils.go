package internal

import (
	"crypto/md5"
	"encoding/base64"
)

func ShortUrlGenerator(url string) string {
	data := []byte(url)

	hasher := md5.New()
	hasher.Write(data)

	c := hasher.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(c)
}
