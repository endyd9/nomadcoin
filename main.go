package main

import (
	"github.com/endyd9/nomadcoin/cli"
	"github.com/endyd9/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
