package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/endyd9/nomadcoin/explorer"
	"github.com/endyd9/nomadcoin/rest"
)

func usage() {
	fmt.Printf("\nWelcome to NomadCoin\n\n")
	fmt.Printf("Please use the following commands:\n")
	fmt.Printf("go run main.go 'mode': Chose between 'html' and 'rest'\n")
	fmt.Printf("-port=portNum: Set the PORT of the Server it's can skipped\n\n")
	fmt.Printf("or if you want run both servers choose 'both' then port is -restPort=portNum -htmlPort=portNum \n\n")
	runtime.Goexit()
}

func parsePort() (int, int) {
	flsgSet := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	if os.Args[1] == "both" {
		port := flsgSet.Int("restPort", 4000, "Set port of the Server")
		port2 := flsgSet.Int("htmlPort", 3000, "Set port of the Server")
		flsgSet.Parse(os.Args[2:])
		if *port == *port2 {
			if *port != 3000 {
				return *port, 3000
			} else if *port2 != 4000 {
				return 4000, *port2
			}
		}
		return *port, *port2
	}
	port := flsgSet.Int("port", 4000, "Set port of the Server")
	flsgSet.Parse(os.Args[2:])
	return *port, 0
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port, port2 := parsePort()

	switch os.Args[1] {
	case "rest":
		rest.Start(port)
	case "html":
		explorer.Start(port)
	case "both":
		go explorer.Start(port2)
		rest.Start(port)
	default:
		usage()
	}
}
