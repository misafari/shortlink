package internal

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func ShortUrlGenerator(key, url string) string {
	data := []byte(fmt.Sprintf("%s:%s", key, url))

	hasher := md5.New()
	hasher.Write(data)

	c := hasher.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(c)
}
