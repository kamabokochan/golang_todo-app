package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config/config.ini") // 実行する場所からの相対パス
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		// MustString: https://pkg.go.dev/gopkg.in/ini.v1#Key.MustString
		// キー値が空の場合、MustString はデフォルト値を返します。
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
