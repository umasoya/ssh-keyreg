package main

import (
	"fmt"
	"flag"

	"github.com/yasuto777/ssh-keyreg/pkg"
)

const (
	token = ""
)

func main() {
	var (
		key_type     = flag.String("t", "rsa", "key type")
		key_len      = flag.Int("b", 4096, "key length")
		filename     = flag.String("f", "github", "filename")
	)
	flag.Parse()

	pkg.GenerateKey(*key_type, fmt.Sprint(*key_len), *filename)
	pub_key := pkg.ReadPublicKey(*filename)

	fmt.Println(pub_key)
}
