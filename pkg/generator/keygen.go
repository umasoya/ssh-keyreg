package generator

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/yasuto777/ssh-keyreg/pkg/options"
)

func Keygen(opts options.Options) {
	var (
		stdout,stderr bytes.Buffer
	)

	cmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", "4096", "-f", opts.KeyPath, "-C", opts.Comment)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		panic(fmt.Sprint(err) + ":" + stderr.String())
	}

	err = cmd.Wait()
	if err != nil {
		if stderr.String() == "" {
			fmt.Println("Skip generate key.(Already key is exists.)")
			return
		} else {
			panic(fmt.Sprint(err) + ":" + stderr.String())
		}
	}
	fmt.Println("Generate " + opts.KeyPath + " is successed.")
	return
}
