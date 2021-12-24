package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: invalide argument")
		os.Exit(1)
	}
	fmt.Println(args)
}

