package main

import (
	"github.com/zzingobomi/zzingocoin/cli"
	"github.com/zzingobomi/zzingocoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
