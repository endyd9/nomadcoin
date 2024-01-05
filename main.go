package main

import (
	"github.com/endyd9/nomadcoin/explorer"
	"github.com/endyd9/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
