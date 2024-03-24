package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Take the data from the block

// Create a counter (nonce) which starts at 0

// Create a hash of the data plus the counter

// Check the hash to see if it meets the set of requirements

// Requirements
// First few bytes of the hash must contain zeros

const Difficulty = 18

// we are setting the difficulty to 18
// the higher the difficulty, the longer it takes to mine a block

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// we are creating a new big int with the value of 1
	target.Lsh(target, uint(256-Difficulty))
	// we are shifting the target to the left by (256 - Difficulty) bits

	pow := &ProofOfWork{b, target}
	// we are creating a new proof of work with the block and the target

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	// we are joining the previous hash, the data, the nonce, and the difficulty
	// we are returning the data

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		// we are creating a hash of the data
		fmt.Printf("\r%x", hash)
		// we are printing the hash
		intHash.SetBytes(hash[:])
		// we are setting the int hash to the hash
		if intHash.Cmp(pow.Target) == -1 {
			break
			// if the int hash is less than the target, we break
		} else {
			nonce++
			// if the int hash is not less than the target, we increment the nonce
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)
	// we are getting the data
	hash := sha256.Sum256(data)
	// we are hashing the data
	intHash.SetBytes(hash[:])
	// we are setting the int hash to the hash

	return intHash.Cmp(pow.Target) == -1
	// we are returning if the int hash is less than the target
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}
