// Package imgconv implements jpeg to png image converter
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

type myImage image.Image

func writeImage(file io.Writer, img myImage) error {
	err := png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}

func readImage(file io.Reader) (img myImage, err error) {
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

// JpgToPng converts jpeg format image files in the directory passed as command line argument to png format image files.
// Even if the specified directory has subdirectories, image files under the subdirectories will be converted.
// If no directory is passed as an argument, an error is returned.
// If more than one directory is passed, it will search the directories in the order they are passed.
// Even if a text file or other file not to be converted is found during the search, it will continue to convert other files.
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
			if !(strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg")) {
				fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
				return nil
			}
			in_path := path
			out_path := replaceSuffix(in_path, getFileExtensionFromFilePath(in_path), ".png")
			if err := convertImage(in_path, out_path); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", trimError(err))
				return nil
			}
			return nil
		})
	}
	return nil
}

func getFileExtensionFromFilePath(filePath string) string {
	dot_at := 0
	for i := len(filePath) - 1; i > -1; i-- {
		if filePath[i] == '.' {
			dot_at = i
			break
		}
	}
	return filePath[dot_at:]
}
