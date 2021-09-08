package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	openKattisPath = "https://open.kattis.com/problems/"
)

func main() {
	fmt.Println(GetProblemById("cold"))
}

func GetProblemById(id string) string {
	resp, err := http.Get(openKattisPath + id)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
