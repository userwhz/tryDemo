package main

import (
	"flag"
	"fmt"
)

func main() {
	pipeName := flag.String("pipe", "", "Specify the name of the named pipe")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *pipeName)
}
