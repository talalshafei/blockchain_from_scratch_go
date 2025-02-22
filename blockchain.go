package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	last_index := len(bc.blocks) - 1
	prevBlock := bc.blocks[last_index]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
