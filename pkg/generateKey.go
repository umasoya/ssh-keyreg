package pkg

import (
	"os/user"
	"os/exec"
	"bytes"
	"fmt"
)

func GenerateKey(t string, l string, f string, m string) {
	var (
		out bytes.Buffer
		stderr bytes.Buffer
	)
	usr, _ := user.Current()
	sshDir := usr.HomeDir + "/.ssh"


	cmd := exec.Command("ssh-keygen", "-t", t, "-b", l, "-f", sshDir + "/" + f, "-C", m)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		panic(fmt.Sprint(err) + ":" + stderr.String())
	}

	err = cmd.Wait()
	if err != nil {
		if stderr.String() == "" {
			fmt.Println("key is already exists")
			return
		} else {
			panic(fmt.Sprint(err) + ":" + stderr.String())
		}
	}
	fmt.Println("Create key.")
}
