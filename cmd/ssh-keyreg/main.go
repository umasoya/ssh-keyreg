package main

import (
	"fmt"
	"log"

	keyreg "github.com/yasuto777/ssh-keyreg"
	"github.com/yasuto777/ssh-keyreg/pkg/options"
)

const version = "2.0"

func main() {
	opts := options.Parse(version)
	err := keyreg.Run(opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successed!")
}
