package main

import (
	"github.com/ReddyyZ/urlbrute/cli"
	"github.com/ReddyyZ/urlbrute/core"
)

func main() {
	config := core.NewConfig("1.0.0", "Tool for brute-force directories on websites")
	cli.Run(config)
}
