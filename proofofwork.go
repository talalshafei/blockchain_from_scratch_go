package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const (
	targetBits = 24
	maxNonce   = math.MaxInt64
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{b, target}
}

func (p *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			p.block.PrevBlockHash,
			p.block.HashTransactions(),
			IntToHex(p.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (p *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining a new block")

	for nonce < maxNonce {
		data := p.prepareData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(p.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}

func (p *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(p.target) == -1

	return isValid
}
