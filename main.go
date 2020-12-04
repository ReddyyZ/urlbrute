package main

import (
	"urlbrute/cli"
	"urlbrute/core"
)

func main() {
	config := core.NewConfig("1.0.0", "Tool for brute-force directories on websites")
	cli.Run(config)
}
