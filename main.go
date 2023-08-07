package main

import (
	"github.com/zzingobomi/zzingocoin/blockchain"
	"github.com/zzingobomi/zzingocoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
