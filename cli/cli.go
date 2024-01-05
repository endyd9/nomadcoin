package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/endyd9/nomadcoin/explorer"
	"github.com/endyd9/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to NomadCoin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-port=4000: Set the PORT of the Server\n")
	fmt.Printf("-mode=rest: Chose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the Server")
	mode := flag.String("mode", "rest", "Chose between 'html' and 'rest'")

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
