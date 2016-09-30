package main

import (
	"encoding/json"
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/hpcloud/tail"
	"io/ioutil"
	"log"
)

type Config struct {
	Target []string `json:"target"`
	Log    string   `json:"log"`
	Port   int      `json:"port"`
}

var config_json string

func init() {
	flag.StringVar(&config_json, "f", "config.json", "config file in json format.")
}

func main() {
	flag.Parse()
	content, _ := ioutil.ReadFile(config_json)
	var config Config
	json.Unmarshal(content, &config)

	go tailFile(config.Log)
	watchFiles(config.Target)
}

func tailFile(f string) {
	t, _ := tail.TailFile(f, tail.Config{Follow: true})
	for line := range t.Lines {
		log.Println(line.Text)
	}
}

func watchFiles(watch_paths []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	for _, f := range watch_paths {
		err = watcher.Add(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	<-done
}