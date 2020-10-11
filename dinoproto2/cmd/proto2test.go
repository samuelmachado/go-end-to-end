package main

import (
	"flag"
	"strings"
)

func main() {
	op := flag.String("op", "s", "s for server", "c for client, and")
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto2Server()

	case "c":
		RunProto2Client()
	}

}
