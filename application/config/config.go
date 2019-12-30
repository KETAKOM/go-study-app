package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	GoogleSpreadSheet GoogleSpreadSheetConfig
}

//Config 設定ファイル
type GoogleSpreadSheetConfig struct {
	AUTH_DIR   string
	SHEET_ID   string
	SHEET_NAME string
}

func LoadConfig() (Config, error) {
	var conf Config
	dir := os.Getenv("CONFIG_DIR")
	_, err := toml.DecodeFile(dir, &conf)
	if err != nil {
		fmt.Println("shippai")
		return Config{}, err
	}
	return conf, nil
}
