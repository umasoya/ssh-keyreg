package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/yasuto777/ssh-keyreg/pkg/options"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Github    Secret
	Bitbucket Secret
}

type Secret struct {
	User  string
	Token string
}

func RegistKey(opts options.Options) {
	key := getPublicKey(opts)
	postApi(opts, key)
	appendHost(opts)
}

func postApi(opts options.Options, key string) {
	conf := getConf(opts)
	req := new(http.Request)

	switch opts.Service {
	case "github":
		req = github(key, conf)
	case "bitbucket":
		req = bitbucket(key, conf)
	default:
		panic("Invalid service.")
	}

	client := new(http.Client)
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return
}

func github(key string, conf Config) *http.Request {
	url := "https://api.github.com/user/keys"
	method := "POST"
	title := "ssh-keyreg"

	jsonStr := `{"title":"` + title + `", "key":"` + key + `"}`
	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Authorization", "token " + conf.Github.Token)

	return req
}

func bitbucket(key string, conf Config) *http.Request {
	url := "https://bitbucket.org/users" + conf.Bitbucket.User + "/ssh-keys"
	method := "POST"
	// title := "ssh-keyreg"

	jsonStr := ``
	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(jsonStr)))
	// req.Header.Set()

	return req
}

/**
* hoge.pubの中身取得
* @param options.Options
* @return string
*/
func getPublicKey(opts options.Options) string {
	fp, err := os.Open(opts.KeyPath + ".pub")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	return scanner.Text()
}

/**
* tomlファイルをパース
*/
func getConf(opts options.Options) Config {
	var conf Config
	_, err := toml.DecodeFile(opts.Token, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}

/**
* configファイルにホストを追加
*/
func appendHost(opts options.Options) {
	if isDuplicate(opts) {
		// 重複してればskip
		return
	}
	// config に追記
	writeConfig(opts)
}

/**
* 重複チェック
*/
func isDuplicate(opts options.Options) bool {
	fp, err := os.Open(opts.Config)
	if err != nil {
		return false
	}
	defer fp.Close()

	lines := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	service := getServiceName(opts)
	if service == "" {
		panic("Invalid service")
	}

	for _,line := range lines {
		if line == "Host " + service {
			return true
		}
	}
	return false
}

/**
* configを新規作成/追記
*/
func writeConfig(opts options.Options) {
	fp, err := os.OpenFile(opts.Config, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	service := getServiceName(opts)
	_, err = fmt.Fprintf(fp, "Host %s\nHostName %s\nIdentityFile %s\nUser git\n", service, service, opts.KeyPath)
	if err != nil {
		panic(err)
	}
}

func getServiceName(opts options.Options) string {
	var service string
	switch opts.Service {
	case "github":
		service = "github.com"
	case "bitbucket":
		service = "bitbucket.org"
	default:
		service = ""
	}
	return service
}
