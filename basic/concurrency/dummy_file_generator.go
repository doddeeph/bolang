package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// https://dasarpemrogramangolang.novalagung.com/A-concurrency-pipeline.html
func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")

	// echo $TMPDIR
	// /var/folders/cr/rv5d85tx0q71s7dkj_3x0nqc0000gp/T/
	// ls -l $TMPDIR/go-concurrency-pipeline-temp
}

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TMPDIR"), "go-concurrency-pipeline-temp") // OSX

func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		content := randomString(contentLength)
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		//sum := fmt.Sprintf("%x", md5.Sum([]byte(content)))
		//filename := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", sum))
		err := ioutil.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("error writing file", filename)
		}

		if i%100 == 0 && i > 0 {
			log.Println(i, "files created")
		}
	}

	log.Printf("%d of total files created", totalFile)
}
