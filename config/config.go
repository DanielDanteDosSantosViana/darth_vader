package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type directory struct {
	Path string `toml:"path"`
}
type config struct {
	Directory directory `toml:"directory"`
}

var Conf config

func Load(config string) {
	if fileInfo, err := os.Stat(config); err != nil {
		if os.IsNotExist(err) {
			log.Panicf("Arquivo de configuração %v não existe.", config)
		} else {
			log.Panicf("Arquivo de configuração %v não pode iniciar. %v", config, err)
		}
	} else {
		if fileInfo.IsDir() {
			log.Panicf("%v é um diretório ", config)
		}
	}

	content, err := ioutil.ReadFile(config)
	if err != nil {
		log.Panicf("read configuration file error. %v", err)
	}
	content = bytes.TrimSpace(content)
	if err := toml.Unmarshal(content, &Conf); err != nil {
		log.Panicf("Erro falta ao tentar carregar o arquivo de configuração. %v", err)
	}
}
