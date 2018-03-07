package pkg

import (
	"bufio"
	"os"
	"os/user"
)

func ReadPublicKey(filename string) string{
	usr, _ := user.Current()
	sshDir := usr.HomeDir + "/.ssh"

	fp, err := os.Open(sshDir + "/" + filename + ".pub")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	scanner.Scan()
	pub_key := scanner.Text()
	return pub_key
}
