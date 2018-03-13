package parser

import (
	"fmt"
	"os"
	"os/user"
	// "path/filepath"
	"log"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Comment   string `short:"C" long:"comment" description:"your_mail@example.com"`
	KeyPath   string `short:"k" long:"key" description:"Path to ssh secret key"`
	Service   []string `short:"s" long:"service" description:"select service to regist(Multiple)" choice:"github" choice:"bitbucket" choice:"gitlab"`
	SshConfig string `short:"F" description:"Path to ssh config file"`
	Toml      string `short:"t" long:"toml" description:"Path to toml file"`
	Version   func() `short:"v" long:"version" description:"show version"`
}

func Parse(ver string) Options {
	opts := Options{Version: func() {
		fmt.Println(ver)
		os.Exit(0)
	}}

	_,err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	setDefaults(&opts)

	return opts
}

func setDefaults(opts *Options) {
	usr,_ := user.Current()
	home := usr.HomeDir

	// toml
	if opts.Toml == "" {
		opts.Toml = "/home/vagrant/Golang/src/github.com/yasuto777/ssh-keyreg/config.toml"
		//exe,_ := os.Executable()
		//opts.Toml = filepath.Dir(exe) + "/config.toml"
	}

	// key path
	if opts.KeyPath == "" {
		opts.KeyPath = home + "/.ssh"
	}

	// ssh config
	if opts.SshConfig == "" {
		opts.SshConfig = home + "/.ssh/config"
	}
}
