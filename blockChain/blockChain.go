package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
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
	return GetBlockChain().blocks[totalBLocks-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockChain().blocks) + 1}
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

var ErrNotFound = errors.New("block not found")

func (b *blockChain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}
	return b.blocks[height-1], nil
}
