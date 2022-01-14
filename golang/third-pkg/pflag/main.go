package main

import (
	"fmt"
	"pflag-example/options"
)

func main() {
	options.Init()
	options.Parse()
	fmt.Println(options.NamedOptions)
}