package main

import (
	"context"
	_ "embed"
	"os"

	"github.com/vizn3r/cloud/cli/cli"
	"github.com/vizn3r/cloud/lib/logger"
)

var log = logger.New("CLI", logger.Green)

func main() {
	log.SetPrintTime(false)
	if err := cli.Commands().Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	log.Close()
}
