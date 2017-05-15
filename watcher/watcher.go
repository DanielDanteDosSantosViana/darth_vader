package watcher

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/DanielDanteDosSantosViana/darth_vader/aws"
	"github.com/DanielDanteDosSantosViana/darth_vader/command"
	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"github.com/DanielDanteDosSantosViana/darth_vader/models"
	"github.com/fsnotify/fsnotify"
)

func Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	clientAWS := aws.NewClientAWS()

	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					filters := config.Conf.Filters
					log.Print(filters)
					for _, filterConf := range filters {
						if filepath.Ext(strings.TrimSpace(event.Name)) == filterConf.Type {
							filter := models.NewFilter(filterConf.S3, filterConf.Email, filterConf.Status, filterConf.TemplateMail, filterConf.Type, event.Name)
							command := command.NewCommand(clientAWS)
							go command.Exec(filter)
						}
					}
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	directories := config.Conf.Directories
	for _, directory := range directories {
		err = watcher.Add(directory.Path)
		if err != nil {
			log.Fatal(err)
		}
	}
	<-done
}
