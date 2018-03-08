package pkg

import (
	"bufio"
	"bytes"
	"net/http"
	"os"
	"os/user"
)

func AddPublicKey(public_key string) error{
	uri    := "https://api.github.com/user/keys"
	method := "POST"
	token := getToken()
	title := "ssh-keyreg"

	jsonStr := `{"title":"` + title + `", "key":"` + public_key + `"}`

	req, err := http.NewRequest(method, uri, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Authorization", "token " + token)

	client := new(http.Client)
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func getToken() string{
	usr, _ := user.Current()
	dir    := usr.HomeDir + "/dotfiles/.local/token"

	fp, err := os.Open(dir + "/github_create_public_key")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	return scanner.Text()
}
