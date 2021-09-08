package testing

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
)

func ParseTestData() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	reader, err := zip.OpenReader(wd + "/cold/samples.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	for _, file := range reader.File {
		fmt.Println(file.Name)
	}
}
