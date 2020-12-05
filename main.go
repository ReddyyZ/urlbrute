package main

import (
	"github.com/ReddyyZ/urlbrute/cli"
	"github.com/ReddyyZ/urlbrute/core"
)

func main() {
	config := core.NewConfig("1.0.2", "Tool for brute-force directories/dns on websites")
	cli.Run(config)
}
