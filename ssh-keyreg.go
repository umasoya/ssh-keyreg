package sshKeyReg

import (
	"fmt"
	// "os"

	"github.com/yasuto777/ssh-keyreg/pkg/options"
	"github.com/yasuto777/ssh-keyreg/pkg/generator"
)

/**
* Host: github.com|bitbucket.org
* PubKey: ~/.ssh/github.pub etc
*/
//type Conf struct {
//	Service string
//	PubKey  *File
//	Token   string
//	Comment string
//	Config  *File
//}

func Run(opts options.Options) {
	generator.Keygen(opts)
}
