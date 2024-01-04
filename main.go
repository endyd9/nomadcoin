package main

import (
	"fmt"

	blockchain "github.com/endyd9/nomadcoin/blockChain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlcok("Second Block")
	chain.AddBlcok("Third Block")
	chain.AddBlcok("Fourth Block")
	for _,block := range chain.AllBlcok() {
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %s\n", block.Hash)
		fmt.Printf("PrevHash : %s\n", block.PrevHash)
 }
}