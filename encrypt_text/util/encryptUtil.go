package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/alwindoss/morse"
	"strings"
)

func EncryptMorse(texto string) string {

	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader("Convert this to Morse"))
	if err != nil {
		return ""
	}
	return string(morseCode)
}

func EncryptMd5(texto string) string {

	hash := md5.Sum([]byte(texto))
	return hex.EncodeToString(hash[:])
}

func EncryptSha512(texto string) string {

	return "Mock ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
}

func EncryptMurcielago(texto string) string {

	return "Mock 01234589"
}
