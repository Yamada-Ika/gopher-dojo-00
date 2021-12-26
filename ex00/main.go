package main

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: invalide argument")
		os.Exit(0)
	}
	for _, dir := range args {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s: no such file or directory\n", dir)
				return err
			}
			fmt.Println(path);
			return nil
		})
	}
}

