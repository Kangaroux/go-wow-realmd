package srpv2

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"math/big"
	"strings"
)

var (
	LargeSafePrime = []byte{
		0x89, 0x4B, 0x64, 0x5E, 0x89, 0xE1, 0x53, 0x5B,
		0xBD, 0xAD, 0x5B, 0x8B, 0x29, 0x06, 0x50, 0x53,
		0x08, 0x01, 0xB1, 0x8E, 0xBF, 0xBF, 0x5E, 0x8F,
		0xAB, 0x3C, 0x82, 0x87, 0x2A, 0x3E, 0x9B, 0xB7,
	}

	xorHash = []byte{
		0xDD, 0x7B, 0xB0, 0x3A, 0x38, 0xAC, 0x73, 0x11,
		0x03, 0x98, 0x7C, 0x5A, 0x50, 0x6F, 0xCA, 0x96,
		0x6C, 0x7B, 0xC2, 0xA7,
	}

	n = toInt(LargeSafePrime)
	g = big.NewInt(7)
	k = big.NewInt(3)
)

// maybe OK (shadowburn). srp6 impl is a bit hidden, but it doesn't seem like this gets reversed
// is 19 bytes *actually* necessary?
func NewPrivateKey() []byte {
	key := make([]byte, 19)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return key
}

// OK (shadowburn)
func CalculateX(username, password string, salt []byte) []byte {
	h := sha1.New()
	inner := sha1.Sum([]byte(strings.ToUpper(username) + ":" + strings.ToUpper(password)))
	h.Write(salt)
	h.Write(inner[:])
	return h.Sum(nil)
}

// OK (shadowburn)
func CalculateVerifier(username, password string, salt []byte) []byte {
	x := big.NewInt(0).SetBytes(Reverse(CalculateX(username, password, salt)))
	return Reverse(pad(32, big.NewInt(0).Exp(g, x, n).Bytes()))
}

// OK (shadowburn)
func CalculateServerPublicKey(verifier []byte, serverPrivateKey []byte) []byte {
	publicKey := big.NewInt(0).Exp(g, toInt(Reverse(serverPrivateKey)), n)
	kv := big.NewInt(0).Mul(k, toInt(Reverse(verifier)))
	return Reverse(pad(32, publicKey.Add(publicKey, kv).Mod(publicKey, n).Bytes()))
}

// OK (shadowburn)
func CalculateU(clientPublicKey, serverPublicKey []byte) []byte {
	h := sha1.New()
	h.Write(clientPublicKey)
	h.Write(serverPublicKey)
	return h.Sum(nil)
}

// OK (shadowburn)
func CalculateServerSKey(clientPublicKey, verifier, u, serverPrivateKey []byte) []byte {
	fmt.Println("---")
	fmt.Printf("A: %x\n", Reverse(clientPublicKey))
	fmt.Printf("v: %x\n", Reverse(verifier))
	fmt.Printf("u: %x\n", Reverse(u))
	fmt.Printf("b: %x\n", Reverse(serverPrivateKey))
	S := big.NewInt(0).Exp(toInt(Reverse(verifier)), toInt(Reverse(u)), n)
	S.Mul(S, toInt(Reverse(clientPublicKey)))
	S.Exp(S, toInt(Reverse(serverPrivateKey)), n)
	return Reverse(pad(32, S.Bytes()))
}

// OK (multiple sources). Shadowburn doesn't drop bytes like it should
func CalculateInterleave(S []byte) []byte {
	for len(S) > 0 && S[0] == 0 {
		S = S[2:]
	}

	lenS := len(S)
	even, odd := make([]byte, lenS/2), make([]byte, lenS/2)

	for i := 0; i < lenS/2; i++ {
		even[i] = S[i*2]
		odd[i] = S[i*2+1]
	}

	hEven := sha1.Sum(even)
	hOdd := sha1.Sum(odd)
	interleaved := make([]byte, 40)

	for i := 0; i < 20; i++ {
		interleaved[i*2] = hEven[i]
		interleaved[i*2+1] = hOdd[i]
	}

	return interleaved
}

// OK (shadowburn)
func CalculateServerSessionKey(clientPublicKey, serverPublicKey, serverPrivateKey, verifier []byte) []byte {
	u := CalculateU(clientPublicKey, serverPublicKey)
	fmt.Printf("u: %x\n", u)
	S := CalculateServerSKey(clientPublicKey, verifier, u, serverPrivateKey)
	fmt.Printf("S: %x\n", S)
	return CalculateInterleave(S)
}

// OK (shadowburn)
func CalculateClientProof(
	username string,
	salt,
	clientPublicKey,
	serverPublicKey,
	sessionKey []byte,
) []byte {
	hUsername := sha1.Sum([]byte(strings.ToUpper(username)))
	h := sha1.New()
	h.Write(xorHash)
	h.Write(hUsername[:])
	h.Write(salt)
	h.Write(clientPublicKey)
	h.Write(serverPublicKey)
	h.Write(sessionKey)
	return h.Sum(nil)
}
