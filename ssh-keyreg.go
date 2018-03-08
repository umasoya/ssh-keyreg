package main

import (
	"fmt"
	"flag"

	"github.com/yasuto777/ssh-keyreg/pkg"
)

func main() {
	var (
		key_type     = flag.String("t", "rsa", "key type")
		key_len      = flag.Int("b", 4096, "key length")
		filename     = flag.String("f", "github", "filename")
		mailAddr	 = flag.String("C", "", "mail address")
	)
	flag.Parse()

	if *mailAddr == "" {
		fmt.Print("Input your email address: ")
		fmt.Scanln(mailAddr)
	}

	pkg.GenerateKey(*key_type, fmt.Sprint(*key_len), *filename, *mailAddr)

	pub_key := pkg.ReadPublicKey(*filename)
	err := pkg.AddPublicKey(pub_key)
	if err != nil {
		panic(err)
	}
	fmt.Println("Create a public key is Successed!")

	err = pkg.RegistClientKey(*filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("All Successed!")
}
