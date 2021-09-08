package main

import (
	"os"

	"chkg.com/kattiscli/problem"
	"chkg.com/kattiscli/testing"
)

func main() {
	switch os.Args[1] {
	case "init":
		problem.InitProblem(os.Args[2])
	case "test":
		testing.ParseTestData()
	}
}
