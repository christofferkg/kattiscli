package problem

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	OpenKattisPath = "https://open.kattis.com/problems/"
	SamplesPath    = "file/statement/samples.zip"
)

func InitProblem(id string) {
	createDirectory(id)
	createWorkFile(id)
	downloadSamples(id)

}

func createDirectory(id string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// Set the pemissions to rwe with 0755
	os.Mkdir(wd+"/"+id, 0755)
}

func createWorkFile(id string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = os.Create(wd + "/" + id + "/main.go")
	if err != nil {
		log.Fatalln(err)
	}
}

func downloadSamples(id string) {
	resp, err := http.Get(OpenKattisPath + id + "/" + SamplesPath)
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
