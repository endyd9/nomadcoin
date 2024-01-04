package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data string
	Hash string
	PrevHash string
}

type blockChain struct {
	blocks []*Block
}


var b *blockChain
var once sync.Once


func (b *Block) getHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBLocks := len(GetBlockChain().blocks)
	if totalBLocks == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBLocks - 1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.getHash()
	return &newBlock
}

func (b *blockChain) AddBlcok(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockChain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{}
			b.AddBlcok("Genesis Block")
		})
	}
	return b
}

func (b *blockChain) AllBlcok() []*Block {
	return b.blocks
}