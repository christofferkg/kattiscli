package problem

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	OpenKattisPath = "https://open.kattis.com/problems/"
	SamplesPath    = "file/statement/samples.zip"
)

var (
	Language = "go"
)

// Get problem creates a directory for the problem, a main file for the
// solution, and downloads the samples into a samples.zip
func GetProblem(id string) {
	mkProblemSpace(id)
	mkSamples(id)
}

func mkDir(name string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Set the pemissions to read-write-execute with 0755
	err = os.Mkdir(wd+"/"+name, 0755)
	if err != nil {
		return err
	}

	return nil
}

func mkFile(name string, extension string) (*os.File, error) {
	file, err := os.Create(fmt.Sprintf("%s.%s", name, extension))
	if err != nil {
		return nil, err
	}

	return file, nil
}

// mkProblemSpace prepares a diretory for the problem by creating a main file,
func mkProblemSpace(id string) error {
	err := mkDir(id)
	if err != nil {
		return err
	}

	err = os.Chdir(id)
	if err != nil {
		return err
	}

	_, err = mkFile("main", Language)
	if err != nil {
		return err
	}

	return nil
}

func DownloadSamples(id string) io.ReadCloser {
	resp, err := http.Get(fmt.Sprintf("%s%s/%s", OpenKattisPath, id, SamplesPath))
	if err != nil {
		log.Fatalf("Failed to download samples: %v", err)
	}
	return resp.Body
}

func mkSamples(id string) {
	samples := DownloadSamples(id)
	defer samples.Close()

	out, err := mkFile("samples", "zip")
	if err != nil {
		log.Fatalf("Failed to create the 'samples.zip': %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, samples)
	if err != nil {
		log.Fatalf("Failed to write samples to 'samples.zip': %v", err)
	}
}
