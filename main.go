package main

import (
	"os"

	"github.com/sophiabrandt/go-embed-youtube/embedyoutube"
)

func main() {
	os.Exit(embedyoutube.CLI(os.Args[1:]))
}
