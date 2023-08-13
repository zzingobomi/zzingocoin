package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/zzingobomi/zzingocoin/explorer"
	"github.com/zzingobomi/zzingocoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 찡오 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:				Set the PORT of the server\n\n")
	fmt.Printf("-mode:				Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set the PORT of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
