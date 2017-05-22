package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"github.com/DanielDanteDosSantosViana/darth_vader/server"
	"github.com/DanielDanteDosSantosViana/darth_vader/watcher"
)

var ID string = os.Getenv("ID_AWS")
var SecretKey string = os.Getenv("SECRET_KEY_AWS")

func init() {
	err := verifyVariableEnv()
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	configFile := flag.String("config", "conf.toml", "Path para o arquivo de configuração")
	flag.Parse()
	if *configFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	config.Load(*configFile)
	config.Conf.Credentials.Id = ID
	config.Conf.Credentials.SecretKey = SecretKey
	server := server.NewServer()
	go server.Listen(":3000")
	watcher.Watch()

}

func verifyVariableEnv() error {
	if ID == "" || SecretKey == "" {
		return fmt.Errorf("Não foi encontrado as variáveis de ambiente para autenticação na AWS 'ID_AWS' e 'SECRET_KEY_AWS'.")
	}
	return nil
}
