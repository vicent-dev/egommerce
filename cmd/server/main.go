package main

import (
	"fmt"
	"os"

	"github.com/vicent-dev/egommerce/pkg/http"
)

func main() {
	if err := http.InitServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
