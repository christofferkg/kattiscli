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
	DownloadSamples("cold")
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

	out, err := os.Create(wd + "/samples.zip")
	if err != nil {
		log.Fatalln(err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}
