package watcher

import (
	"log"

	"github.com/DanielDanteDosSantosViana/darth_vader/aws"
	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"github.com/DanielDanteDosSantosViana/darth_vader/reader"
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
					go reader.NewReader(event.Name, clientAWS).Read()
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
