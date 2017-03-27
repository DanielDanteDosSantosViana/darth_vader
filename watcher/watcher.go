package watcher

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DanielDanteDosSantosViana/darth_vader/aws"
	"github.com/DanielDanteDosSantosViana/darth_vader/config"
	"github.com/fsnotify/fsnotify"
)

func Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	client := aws.NewClientAWS()
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					file, e := os.Open(event.Name)
					if e != nil {
						fmt.Printf("err opening file: %s", e)
					}
					defer file.Close()
					fileInfo, _ := file.Stat()
					size := fileInfo.Size()
					buffer := make([]byte, size)
					file.Read(buffer)
					fileBytes := bytes.NewReader(buffer)
					fileType := http.DetectContentType(buffer)
					path := event.Name
					client.SendToS3(aws.NewParams(config.Conf.Bucket.Name, path, size, fileBytes, fileType))
					log.Println("Arquivo novo identificado : ", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(config.Conf.Directory.Path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
