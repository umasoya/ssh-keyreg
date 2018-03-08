package pkg

import (
	"fmt"
	"bufio"
	"os"
	"os/user"
)

func RegistClientKey(filename string) error{
	if isDuplicate() {
		return nil
	}
	// 重複しなければ config に追記
	return writeConfig(filename)
}

// 既に~/.ssh/config に github があれば true
func isDuplicate() bool{
	usr, _ := user.Current()

	fp, err := os.Open(usr.HomeDir + "/.ssh/config")
	// config file is not exists.
	if err != nil {
		return false
	}
	defer fp.Close()

	lines := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// Can't scan file
	if scan_err := scanner.Err(); scan_err != nil {
		panic(err)
	}

	// If exists "Host github.com" in config, return true
	for _, line := range lines {
		if line == "Host github.com" {
			return true
		}
	}
	return false
}

// config に github を追記
func writeConfig(filename string) error{
	usr, _ := user.Current()

	fp, err := os.OpenFile(usr.HomeDir + "/.ssh/config", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()

	host := `Host github.com
	HostName github.com
	IdentityFile ` + usr.HomeDir + "/.ssh/" + filename + `
	User git`

	_, err = fmt.Fprintf(fp, "%s\n", host)
	if err != nil {
		return err
	}
	return nil
}
