package blockchain

import (
	"sync"

	"github.com/endyd9/nomadcoin/db"
	"github.com/endyd9/nomadcoin/utils"
)

type blockChain struct {
	NeweatHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockChain
var once sync.Once

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockIntervla      int = 2
	allowedRange       int = 2
)

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
	b.CurrentDifficulty = block.Difficulty
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

func (b *blockChain) recalculateDifficalty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockIntervla

	if actualTime < (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockChain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.recalculateDifficalty()
	} else {
		return b.CurrentDifficulty
	}
}

func BlockChain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{
				Height: 0,
			}
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
