package main

import (
	"errors"
	"flag"
	"io/fs"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var iFlag = flag.String("i", "jpg", "input file extension")
var oFlag = flag.String("o", "png", "output file extension")

func convert(path string) (err error) {
	if !strings.HasSuffix(path, ".jpg") {
		return errors.New(fmt.Sprintf("error: %s is not a valid file\n", path))
	}
	in_path := path
	out_path := strings.Replace(path, ".jpg", ".png", 1)
	in_file, err := os.Open(in_path)
	if err != nil {
		return err
	}
	in_img, err := jpeg.Decode(in_file)
	if err != nil {
		return err
	}
	defer func() {
		err = in_file.Close()
	}()
	out_file, err := os.Create(out_path)
	if err != nil {
		return err
	}
	if err := png.Encode(out_file, in_img); err != nil {
		out_file.Close()
		return err
	}
	defer func() {
		err = out_file.Close()
	}()
	return nil
}

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
			if err := convert(path); err != nil {
				return errors.New("faile : convert")
			}
			return nil
		})
	}
}
