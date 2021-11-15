package main

import (
	"os"

	"chkg.com/kattiscli/problem"
	"chkg.com/kattiscli/testing"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "kattiscli"

	app.Commands = []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Create a directory and download the test files for a given problem",
			Action: func(c *cli.Context) {
				// Assume that the problem id is always the last os.Arg
				problem.GetProblem(os.Args[len(os.Args)-1])
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Test the main file against the provided sample file",
			Action: func(c *cli.Context) {
				testing.ParseTestData()
			},
		},
	}

	app.Run(os.Args)
}
