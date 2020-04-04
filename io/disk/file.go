package disk

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

func WatchFile(f string, grace time.Duration) chan []byte {
	result := make(chan []byte)

	go func() {
		fileToChan(f, result)

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			logrus.Fatalln(err)
		}
		defer watcher.Close()

		done := make(chan bool)
		go func() {
			for {
				last := time.Now()
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						if time.Since(last) > grace {
							err := fileToChan(f, result)
							if err != nil {
								time.Sleep(grace)
								last = time.Now()
							}
						}
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					fmt.Println("error:", err)
				}
			}
		}()

		err = watcher.Add(f)
		if err != nil {
			logrus.Panicln(err)
		}
		<-done

	}()
	return result
}

func fileToChan(f string, result chan []byte) (err error) {
	var b []byte
	b, err = ioutil.ReadFile(f)
	if err != nil {
		logrus.Errorln(fmt.Errorf("could not read file %v", err))
		return err
	} else {
		result <- b
		logrus.Println("Successfully read config", f)
	}
	return
}
