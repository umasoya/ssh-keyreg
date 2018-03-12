package sshKeyReg

import (
	"github.com/yasuto777/ssh-keyreg/pkg/options"
	"github.com/yasuto777/ssh-keyreg/pkg/generator"
)

func Run(opts options.Options) {
	// キーペアの生成
	generator.Keygen(opts)
	// API投げて公開鍵を登録
	// configにホスト情報を追記
	generator.RegistKey(opts)
}
