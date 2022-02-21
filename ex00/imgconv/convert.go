// Package imgconv implements image converter
package imgconv

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Image image.Image

func writeImage(file io.Writer, img Image) error {
	err := png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}

func readImage(file io.Reader) (img Image, err error) {
	img, err = jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func convertImage(in_path string, out_path string) error {
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

// JpgToPng converts jpeg image file to png image file
func JpgToPng() error {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return errors.New("error: invalid argument")
	}
	for _, dir := range args {
		filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", trimError(err))
				return nil
			}
			if info.IsDir() || strings.HasSuffix(path, ".png") {
				return nil
			}
			if !strings.HasSuffix(path, ".jpg") {
				fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
				return nil
			}
			in_path := path
			out_path := replaceSuffix(path, ".jpg", ".png")
			if err := convertImage(in_path, out_path); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", trimError(err))
				return nil
			}
			return nil
		})
	}
	return nil
}
