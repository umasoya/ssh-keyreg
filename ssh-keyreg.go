package sshKeyReg

import (
	"log"

	"github.com/yasuto777/ssh-keyreg/pkg/options"
	"github.com/yasuto777/ssh-keyreg/pkg/generator"
)

func Run(opts options.Options) error {
	// キーペアの生成
	err := generator.Keygen(opts)
	if err != nil {
		return err
	}
	// API投げて公開鍵を登録
	// configにホスト情報を追記
	err = generator.RegistKey(opts)
	if err != nil {
		return err
	}
	return nil
}
