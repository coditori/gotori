package main

import (
	"flag"
	"fmt"
	"os"
	"rest/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	println("environment is " + *environment)
	server.Init()
}
