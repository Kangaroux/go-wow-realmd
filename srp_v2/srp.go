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

	xorHash = []byte{
		0xDD, 0x7B, 0xB0, 0x3A, 0x38, 0xAC, 0x73, 0x11,
		0x03, 0x98, 0x7C, 0x5A, 0x50, 0x6F, 0xCA, 0x96,
		0x6C, 0x7B, 0xC2, 0xA7,
	}

	n = toInt(largeSafePrime)
	g = big.NewInt(7)
	k = big.NewInt(3)
)

func calculateX(username, password string, salt []byte) []byte {
	h := sha1.New()
	h.Write(salt)
	inner := sha1.Sum([]byte(strings.ToUpper(username) + ":" + strings.ToUpper(password)))
	h.Write(inner[:])
	return reverse(h.Sum(nil))
}

func calculateVerifier(username, password string, salt []byte) []byte {
	x := big.NewInt(0).SetBytes(calculateX(username, password, salt))
	return pad(32, big.NewInt(0).Exp(g, x, n).Bytes())
}

func calculateServerPublicKey(verifier []byte, serverPrivateKey []byte) []byte {
	publicKey := big.NewInt(0).Exp(g, toInt(serverPrivateKey), n)
	kv := big.NewInt(0).Mul(k, toInt(verifier))
	publicKey.Add(publicKey, kv).Mod(publicKey, n)
	return pad(32, publicKey.Bytes())
}

func calculateU(clientPublicKey, serverPublicKey []byte) []byte {
	h := sha1.New()
	h.Write(reverse(clientPublicKey))
	h.Write(reverse(serverPublicKey))
	return reverse(h.Sum(nil))
}

func calculateServerSKey(clientPublicKey, verifier, u, serverPrivateKey []byte) []byte {
	S := big.NewInt(0).Exp(toInt(verifier), toInt(u), n)
	S.Mul(S, toInt(clientPublicKey))
	S.Exp(S, toInt(serverPrivateKey), n)
	return reverse(pad(32, S.Bytes()))
}

func calculateInterleave(S []byte) []byte {
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

func calculateServerSessionKey(clientPublicKey, serverPublicKey, serverPrivateKey, verifier []byte) []byte {
	u := calculateU(clientPublicKey, serverPublicKey)
	S := calculateServerSKey(clientPublicKey, verifier, u, serverPrivateKey)
	return calculateInterleave(S)
}

func calculateClientProof(
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
	h.Write(reverse(clientPublicKey))
	h.Write(reverse(serverPublicKey))
	h.Write(sessionKey)
	return h.Sum(nil)
}
