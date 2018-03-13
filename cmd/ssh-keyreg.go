package main

import (
	// "fmt"
	"log"
	"os"

	"github.com/yasuto777/ssh-keyreg/pkg/parser"
	"github.com/yasuto777/ssh-keyreg/pkg/generator"
)

const version = "2.0"

func main() {
	// option parse
	opts := parser.Parse(version)
	// get toml context
	conf := parser.ParseToml(&opts)

	// fmt.Printf("%#v\n",opts.Service)
	//fmt.Printf("%#v\n",conf)
	//fmt.Printf("%#v\n",conf.Github.Token)

	// generate ssh key
	err := generator.Generate(&opts, &conf)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
