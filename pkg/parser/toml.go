package parser

import (
	"os"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	User  string `toml:"user"`
	Token string `toml:"token"`
	Label string `toml:"label"`
}

type Service struct {
	Github    Config
	Bitbucket Config
	Gitlab    Config
}

func ParseToml(opts *Options) Service {
	var service Service
	_,err := toml.DecodeFile(opts.Toml, &service)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return service
}
