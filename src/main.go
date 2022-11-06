package main

import (
	"flag"

	"github.com/ssibrahimbas/cli-todos/src/app"
)

func main() {
	err := app.New().Init().Start()
	if err != nil {
		flag.PrintDefaults()
	}
}