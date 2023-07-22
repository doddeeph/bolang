package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	log.Println("start")
	start := time.Now()

	//chanFileContent := readFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func readFiles() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		err := filepath.Walk(tempPath2, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			chanOut <- FileInfo{
				Filepath: path,
				Content:  buf,
			}
			return nil
		})
		if err != nil {
			log.Println("error:", err.Error())
		}
		close(chanOut)
	}()
	return chanOut
}

var tempPath2 = filepath.Join(os.Getenv("TEMP"), "go-concurrency-pipeline-temp")

type FileInfo struct {
	Filepath  string
	Content   []byte
	Sum       string
	IsRenamed bool
}
