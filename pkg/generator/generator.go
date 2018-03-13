package generator

import (
	"sync"
	"strings"
	"os/exec"
	"os"
	"fmt"

	"github.com/yasuto777/ssh-keyreg/pkg/parser"
	"golang.org/x/crypto/ssh/terminal"
)

func Generate(opts *parser.Options, conf *parser.Service) error {
	var (
		wg sync.WaitGroup
		options = "-t rsa -b 4096"
	)

	// set pass phrase
	pass := setPassPhrase(opts)
	fmt.Printf("%#v\n",pass)
	os.Exit(0)

	for _,service := range opts.Service {
		wg.Add(1)
		go func(service string, pass map[string]string) {
			if pass["all"] != "" {
				optionStr := options + " -N " + pass["all"] + " -F " + opts.KeyPath + "/" + service
				if opts.Comment != "" {
					optionStr = optionStr + " -C " + opts.Comment
				}
			} else {
				optionStr := options + " -N " + pass[service] + " -F " + opts.KeyPath + "/" + serive
				if opts.Comment != "" {
					optionStr = optionStr + " -C " + opts.Comment
				}
			}
			cmd := exec.Command("ssh-keygen", optionStr)
			wg.Done()
		}(service, pass)
	}

	wg.Wait()
	fmt.Println("finish")
	return nil
}

func setPassPhrase(opts *parser.Options) map[string]string {
	var yn string
	m := map[string]string{}

	fmt.Printf("Use the same passphrase for all? (Y/N): ")
	fmt.Scanln(&yn)
	yn = strings.ToUpper(yn)

	if yn == "Y" {
		fmt.Printf("Please enter passphrase: ")
		pass,_ := terminal.ReadPassword(0)
		m["all"] = fmt.Sprintf("%s",pass)
	} else if yn == "N" {
		// Scan every phrase
		for _,s := range opts.Service {
			fmt.Printf("For %s passphrase: ",s)
			pass,_ := terminal.ReadPassword(0)
			fmt.Printf("\n")
			m[s] = fmt.Sprintf("%s",pass)
		}
	} else {
		panic("panic")
	}

	return m
}
