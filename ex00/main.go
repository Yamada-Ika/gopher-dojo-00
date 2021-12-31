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

func convert(path string) error {
	if !strings.HasSuffix(path, ".jpg") {
		return errors.New(fmt.Sprintf("error: %s is not a valid file\n", path))
	}
	jpg_path := path
	png_path := strings.Replace(path, ".jpg", ".png", 1)
	jpg_file, err := os.Open(jpg_path)
	if err != nil {
		return err
	}
	jpg_img, err := jpeg.Decode(jpg_file)
	if err != nil {
		return err
	}
	if err := jpg_file.Close(); err != nil {
		return err
	}
	png_file, err := os.Create(png_path)
	if err != nil {
		return err
	}
	if err := png.Encode(png_file, jpg_img); err != nil {
		png_file.Close()
		return err
	}
	if err := png_file.Close(); err != nil {
		return err
	}
	return nil
}

var iFlag = flag.String("i", "jpg", "input file extension")
var oFlag = flag.String("o", "png", "output file extension")

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
