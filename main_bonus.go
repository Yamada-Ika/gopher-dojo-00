package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func validateFlag() error {
	switch *inputFileFormat {
	case "jpg", "png", "gif":
		break
	default:
		return errors.New("error: invalide extension")
	}
	switch *outputFileFormat {
	case "jpg", "png", "gif":
		break
	default:
		return errors.New("error: invalide extension")
	}
	return nil
}

var inputFileFormat = flag.String("i", "jpg", "input file extension")
var outputFileFormat = flag.String("o", "png", "output file extension")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: invalid argument")
		return
	}
	if err := validateFlag(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	inputFileExt := "." + *inputFileFormat
	outputFileExt := "." + *outputFileFormat
	for _, dir := range args {
		filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s: no such file or directory\n", path)
				return err
			}
			if info.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, inputFileExt) {
				fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
				return nil
			}
			in_path := path
			out_path := replaceSuffix(path, inputFileExt, outputFileExt)
			if err := convertImage(in_path, out_path); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}
			return nil
		})
	}
}
