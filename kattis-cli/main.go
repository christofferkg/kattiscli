package main

import (
	"os"

	"chkg.com/kattiscli/problem"
)

func main() {
	problem.InitProblem(os.Args[1])
}
