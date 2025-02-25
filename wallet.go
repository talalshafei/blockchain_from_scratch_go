package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

const (
	version            = byte(0x00)
	addressChecksumLen = 4
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)

	checksum := checksum(versionedPayload)
	fullPayload := append(versionedPayload, checksum...)

	address := Base58Encode(fullPayload)

	return address
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.Y.Bytes()...)
	return *private, pubKey
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()

	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}

	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256((firstSHA[:]))

	return secondSHA[:addressChecksumLen]
}

func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	checksumStartIndex := len(pubKeyHash) - addressChecksumLen

	actualChecksum := pubKeyHash[checksumStartIndex:]
	pkVersion := pubKeyHash[0]

	pubKeyHash = pubKeyHash[1:checksumStartIndex]
	targetChecksum := checksum(append([]byte{pkVersion}, pubKeyHash...))

	return bytes.Equal(actualChecksum, targetChecksum)
}

// To fix the error when that was caused by SaveToFile function
// error was: gob: type elliptic.p256Curve has no exported fields, panic: gob: type elliptic.p256Curve has no exported fields
func (w *Wallet) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	// Encode the necessary components
	if err := encoder.Encode(w.PrivateKey.D); err != nil {
		return nil, err
	}
	if err := encoder.Encode(w.PrivateKey.PublicKey.X); err != nil {
		return nil, err
	}
	if err := encoder.Encode(w.PrivateKey.PublicKey.Y); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (w *Wallet) GobDecode(data []byte) error {
	buf := bytes.NewReader(data)
	decoder := gob.NewDecoder(buf)

	var D, X, Y *big.Int

	// Decode the components
	if err := decoder.Decode(&D); err != nil {
		return err
	}
	if err := decoder.Decode(&X); err != nil {
		return err
	}
	if err := decoder.Decode(&Y); err != nil {
		return err
	}

	// Reconstruct the private key using P256 curve
	curve := elliptic.P256()
	w.PrivateKey = ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     X,
			Y:     Y,
		},
		D: D,
	}

	// Reconstruct the public key bytes from X and Y
	w.PublicKey = append(w.PrivateKey.PublicKey.X.Bytes(), w.PrivateKey.PublicKey.Y.Bytes()...)

	return nil
}
