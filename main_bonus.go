package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Image image.Image

func writeImage(file io.Writer, img Image) (err error) {
	switch *oFlag {
	case "jpg":
		err = jpeg.Encode(file, img, nil)
	case "png":
		err = png.Encode(file, img)
	case "gif":
		err = gif.Encode(file, img, nil)
	}
	if err != nil {
		return err
	}
	return nil
}

func readImage(file io.Reader) (img Image, err error) {
	switch *iFlag {
	case "jpg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	}
	if err != nil {
		return nil, err
	}
	return img, nil
}

func convert(in_path string, out_path string) (err error) {
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

func flagValidate() error {
	switch *iFlag {
	case "jpg", "png", "gif":
		return nil
	default:
		return errors.New("error: invalide extension")
	}
	switch *oFlag {
	case "jpg", "png", "gif":
		return nil
	default:
		return errors.New("error: invalide extension")
	}
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
	if err := flagValidate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	iExt := "." + *iFlag
	oExt := "." + *oFlag
	for _, dir := range args {
		filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s: no such file or directory\n", path)
				return err
			}
			if info.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, iExt) {
				fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
				return nil
			}
			in_path := path
			out_path := strings.Replace(path, iExt, oExt, 1)
			if err := convert(in_path, out_path); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}
			return nil
		})
	}
}
