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
		os.Exit(0)
	}
	for _, path := range args {
		if _, err := os.Stat(path); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s: no such file or directory\n", path)
			os.Exit(0)
		}
	}
}

