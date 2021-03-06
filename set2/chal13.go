package set2

import (
	"fmt"
	"log"
	"strings"

	"github.com/sespiros/go-cryptopals/util"
)

func Parse(str string, delim string) (map[string]string, error) {
	vars := strings.Split(str, delim)
	obj := make(map[string]string)

	for _, v := range vars {
		lr := strings.Split(v, "=")
		if len(lr) < 2 {
			return nil, fmt.Errorf("Error in parsing")
		}
		obj[lr[0]] = lr[1]
	}

	return obj, nil
}

func profileFor(email string) string {
	for _, v := range email {
		if v == '=' || v == '&' {
			log.Fatal("profileFor: Invalid characters inserted")
		}
	}

	str := "email="
	str += email
	str += "&uid=10&role=user"

	return str
}

func encryptProfile(email string, key []byte) []byte {
	plain := PKCSPadding([]byte(profileFor(email)), len(key))
	cipher := util.EncryptAESECB(plain, key)
	return cipher
}

func decryptParseProfile(encUserProfile, key []byte) map[string]string {
	plain := util.DecryptAESECB(encUserProfile, key)
	plain = plain[:len(plain)-int(plain[len(plain)-1])] //remove PKCS#7 padding

	profile, err := Parse(string(plain), "&")
	util.Check(err)
	return profile
}
