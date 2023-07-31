package main

import (
	"github.com/zzingobomi/zzingocoin/blockchain"
	"github.com/zzingobomi/zzingocoin/db"
)

func main() {
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
	db.DB().Close()
}
