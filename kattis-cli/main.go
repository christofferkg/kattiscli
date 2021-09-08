package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	openKattisPath = "https://open.kattis.com/problems/"
	samplesPath    = "file/statement/samples.zip"
)

func main() {
	InitProblem(os.Args[1])
}

func InitProblem(id string) {
	CreateDirectory(id)
	CreateWorkFile(id)
	DownloadSamples(id)

}

func CreateWorkFile(id string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = os.Create(wd + "/" + id + "/main.go")
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateDirectory(id string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// Set the pemissions to rwe with 0755
	os.Mkdir(wd+"/"+id, 0755)
}

func DownloadSamples(id string) {
	resp, err := http.Get(openKattisPath + id + "/" + samplesPath)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	out, err := os.Create(wd + "/" + id + "/samples.zip")
	if err != nil {
		log.Fatalln(err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}
