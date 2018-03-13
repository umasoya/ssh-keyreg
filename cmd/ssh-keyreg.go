package main

import (
	"fmt"
	"github.com/yasuto777/ssh-keyreg/pkg/parser"
)

const version = "2.0"

func main() {
	// option parse
	opts := parser.Parse(version)
	// get toml context
	//conf := parser.ParseToml(&opts)

	fmt.Printf("%#v\n",opts)
	//fmt.Printf("%#v\n",conf.Github.Token)
}
