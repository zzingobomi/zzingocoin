package blockchain

import (
	"fmt"
	"sync"

	"github.com/zzingobomi/zzingocoin/db"
	"github.com/zzingobomi/zzingocoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// func (b *blockchain) Blocks() []*Block {}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				fmt.Printf("Restoring...")
				b.restore(checkpoint)
			}
		})
	}
	fmt.Println(b.NewestHash)
	return b
}
