package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// var alphabet = []string{"a", "b", "c"}
var alphabet = "abcdefghijklmnopqrstuvwxyz"
var passwordLength = 8

func dummyHash(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

func decode(hash string, current string) string {
	if len(current) == passwordLength {
		if dummyHash(current) == hash {
			return current
		}
		return ""
	}

	for i := 0; i < len(alphabet); i++ {
		a := current + string(alphabet[i])
		result := decode(hash, a)
		if result != "" {
			return result
		}
	}

	return ""
}

func main() {
	password := "breasddddd"
	hash := dummyHash(password)
	decodedPassword := decode(hash, "")
	fmt.Println(decodedPassword)
}
