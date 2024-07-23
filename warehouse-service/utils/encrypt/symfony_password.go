package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

const (
	// Iteration default in symfony
	Iteration int = 5000
)

// Hash Craete hashing password from plain password and salt
// Need Error Handler, probably panic here
func Hash(password string, salt string) string {
	var digest []byte
	salted := fmt.Sprintf("%s{%s}", password, salt)

	sha512Salted := sha512.New()
	sha512Salted.Write([]byte(salted))
	digest = sha512Salted.Sum(nil)

	for index := 1; index < Iteration; index++ {
		newSHA512 := sha512.New()
		newSHA512.Write([]byte(digest))
		newSHA512.Write([]byte(salted))
		digest = newSHA512.Sum(nil)
	}

	return base64.StdEncoding.EncodeToString(digest)
}

// Verify plain password
func Verify(plain string, salt string, hashed string) bool {

	hashedFromPlain := Hash(plain, salt)

	equal := strings.Compare(hashed, hashedFromPlain)

	return equal == 0
}

// GenerateSalt algorithm by symfony
func GenerateSalt() string {
	const saltLength int = 32
	var saltRegexByteResult []byte
	var regex *regexp.Regexp

	saltByte := make([]byte, saltLength)
	_, err := rand.Read(saltByte)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	saltBase64 := base64.StdEncoding.EncodeToString(saltByte)

	regex = regexp.MustCompile(`\+`)
	// salt = strings.Replace(saltBase64, )
	saltRegexByteResult = regex.ReplaceAll([]byte(saltBase64), []byte("."))

	regex = regexp.MustCompile(`=$`)
	saltRegexByteResult = regex.ReplaceAll(saltRegexByteResult, []byte(""))

	return string(saltRegexByteResult)
}
