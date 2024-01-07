package blockchain

import (
	"sync"

	"github.com/endyd9/nomadcoin/db"
	"github.com/endyd9/nomadcoin/utils"
)

type blockChain struct {
	NeweatHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockChain
var once sync.Once

func (b *blockChain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockChain) persist() {
	db.SaveCheckPoint(utils.ToBytes(b))
}

func (b *blockChain) AddBlcok(data string) {
	block := createBlock(data, b.NeweatHash, b.Height+1)
	b.NeweatHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockChain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NeweatHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func BlockChain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{"", 0}
			checkPoint := db.CheckPoint()
			if checkPoint == nil {
				b.AddBlcok("Genesis Block")
			} else {
				b.restore(checkPoint)
			}
		})
	}
	return b
}
