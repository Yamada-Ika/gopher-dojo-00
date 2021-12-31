package main

import (
	"errors"
	"flag"
	"io"
	"io/fs"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var iFlag = flag.String("i", "jpg", "input file extension")
var oFlag = flag.String("o", "png", "output file extension")

func writeImage(file io.Writer, img image.Image) (error) {
	if err := png.Encode(file, img); err != nil {
		return err
	}
	return nil
}

func readImage(file io.Reader) (img image.Image, err error) {
	img, err = jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func convert(path string) (err error) {
	in_path := path
	out_path := strings.Replace(path, ".jpg", ".png", 1)
	in_file, err := os.Open(in_path)
	if err != nil {
		return err
	}
	in_img, err := readImage(in_file)
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
	if err := writeImage(out_file, in_img); err != nil {
		return err
	}
	defer func() {
		err = out_file.Close()
	}()
	return nil
}

func flagValidate(flag string) (error) {
	switch flag {
	case "jpg", "png":
		return nil
	default:
		return errors.New("error: invalide extension")
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: invalide argument")
		return
	}
	// flag check
	if err := flagValidate(*iFlag); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := flagValidate(*oFlag); err != nil {
		fmt.Fprintln(os.Stderr, err)
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
			if err := convert(path); err != nil {
				return errors.New("error : failed to convert")
			}
			return nil
		})
	}
}
