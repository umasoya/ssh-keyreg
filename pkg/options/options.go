package options

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Service string `long:"service" short:"s" default:"github" choice:"github" choice:"bitbucket" choice:"gitlab" description:"Select the service to register."`
	KeyPath string `long:"path" short:"p" default-mask:"~/.ssh/<service>" description:"Path to save ssh key"`
	Token  string `long:"secret" short:"s" default-mask:"Execution file path" description:"Path of toml file that wrote the username & token."`
	Comment string `short:"C" description:"your_email@example.com"`
	Config  string `short:"c" long:"conf" default-mask:"~/.ssh/config" description:"Path to ssh config file"`
	Version func() `short:"v" long:"version" description:"show version"`
}

func Parse(version string) Options {
	opts := Options{Version: func() {
		fmt.Println(version)
		os.Exit(0)
	}}
	_,err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	setTokenPath(&opts)
	setKeyPath(&opts)
	setComment(&opts)
	setConfig(&opts)

	return opts
}

func getHomeDir() string {
	usr,_ := user.Current()
	return usr.HomeDir
}

func setKeyPath(opts *Options) {
	if opts.KeyPath != "" {
		return
	}
	opts.KeyPath = getHomeDir() + "/.ssh/" + opts.Service
	return
}

func setTokenPath(opts *Options) {
	// return if already set value
	if opts.Token != "" {
		return
	}
	exe,_ := os.Executable()
	opts.Token = filepath.Dir(exe) + "/config.toml"
	return
}

func setComment(opts *Options) {
	if opts.Comment != "" {
		return
	}
	fmt.Print("Plese enter email_address: ")
	fmt.Scanln(&opts.Comment)
	return
}

func setConfig(opts *Options) {
	if opts.Config != "" {
		return
	}
	opts.Config = getHomeDir() + "/.ssh/config"
	return
}
