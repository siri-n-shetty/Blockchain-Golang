package blockchain

type Blockchain struct {
	Blocks []*Block
	// this has an array of pointers to blocks
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// we need to run the proof of work algorith for each block created
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// we are getting the last block in the chain
	new := CreateBlock(data, prevBlock.Hash)
	// we are creating a new block with the data and the hash of the previous block
	chain.Blocks = append(chain.Blocks, new)
	// we are appending the new block to the chain
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
	// this is the first block in the chain
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
	// we are initializing the blockchain with the genesis block
}
