package main

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// https://dasarpemrogramangolang.novalagung.com/A-concurrency-pipeline.html
func main() {
	log.Println("start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

var tempPath1 = filepath.Join(os.Getenv("TEMP"), "go-concurrency-pipeline-temp")

func proceed() {
	counterTotal := 0
	counterRenamed := 0
	err := filepath.Walk(tempPath1, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		counterTotal++
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		sum := fmt.Sprintf("%x", md5.Sum(buf))
		destPath := filepath.Join(tempPath1, fmt.Sprintf("file-%s.txt", sum))
		err = os.Rename(path, destPath)
		if err != nil {
			return err
		}
		counterRenamed++
		return nil
	})
	if err != nil {
		log.Println("error:", err.Error())
	}
	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
}
