package main

import (
	"flag"
	"os"

	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"github.com/DanielDanteDosSantosViana/darth_vader/watcher"
)

func main() {
	configFile := flag.String("config", "conf.toml", "Path para o arquivo de configuração")
	flag.Parse()
	if *configFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	config.Load(*configFile)
	watcher.Watch()

}
