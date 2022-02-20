package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: invalide argument")
		return
	}
	for _, dir := range args {
		filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s: no such file or directory\n", path)
				return err
			}
			if info.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, ".jpg") {
				fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
				return nil
			}
			in_path := path
			out_path := strings.Replace(path, ".jpg", ".png", 1)
			if err := convert(in_path, out_path); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}
			return nil
		})
	}
}
