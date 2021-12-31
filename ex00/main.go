package main

import (
	"flag"
	"io/fs"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
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
			jpg_path := path
			png_path := strings.Replace(path, ".jpg", ".png", 1)
			jpg_file, err := os.Open(jpg_path)
			if err != nil {
				log.Fatal(err)
			}
			jpg_img, err := jpeg.Decode(jpg_file)
			if err != nil {
				log.Fatal(err)
			}
			if err := jpg_file.Close(); err != nil {
				log.Fatal(err)
			}
			png_file, err := os.Create(png_path)
			if err != nil {
				log.Fatal(err)
			}
			if err := png.Encode(png_file, jpg_img); err != nil {
				png_file.Close()
				log.Fatal(err)
			}
			if err := png_file.Close(); err != nil {
				log.Fatal(err)
			}
			return nil
		})
	}
}
