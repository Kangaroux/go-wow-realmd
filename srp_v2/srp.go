package srpv2

import (
	"crypto/sha1"
	"math/big"
	"strings"
)

var (
	largeSafePrime = []byte{
		0x89, 0x4B, 0x64, 0x5E, 0x89, 0xE1, 0x53, 0x5B,
		0xBD, 0xAD, 0x5B, 0x8B, 0x29, 0x06, 0x50, 0x53,
		0x08, 0x01, 0xB1, 0x8E, 0xBF, 0xBF, 0x5E, 0x8F,
		0xAB, 0x3C, 0x82, 0x87, 0x2A, 0x3E, 0x9B, 0xB7,
	}
)

func bigN() *big.Int {
	return big.NewInt(0).SetBytes(largeSafePrime)
}

func bigG() *big.Int {
	return big.NewInt(7)
}

func bigK() *big.Int {
	return big.NewInt(3)
}

func calculateX(username, password string, salt []byte) []byte {
	h := sha1.New()
	h.Write(salt)
	inner := sha1.Sum([]byte(strings.ToUpper(username) + ":" + strings.ToUpper(password)))
	h.Write(inner[:])
	return h.Sum(nil)
}

func calculateVerifier(username, password string, salt []byte) []byte {
	x := big.NewInt(0).SetBytes(reverse(calculateX(username, password, salt)))
	return reverse(big.NewInt(0).Exp(bigG(), x, bigN()).Bytes())
}