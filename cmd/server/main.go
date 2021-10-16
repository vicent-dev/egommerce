package main

import (
	"fmt"
	"os"

	"github.com/vicent-dev/egommerce/app"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {
	return app.NewServer().Serve()
}
