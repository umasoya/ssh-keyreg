package main

import (
	"fmt"

	keyreg "github.com/yasuto777/ssh-keyreg"
	"github.com/yasuto777/ssh-keyreg/pkg/options"
)

const version = "2.0"

func main() {
	opts := options.Parse(version)
	keyreg.Run(opts)
	fmt.Println("Successed!")
}
