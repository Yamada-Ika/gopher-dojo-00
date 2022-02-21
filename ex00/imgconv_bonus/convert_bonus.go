package imgconv_bonus

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
	switch *outputFileFormat {
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
	switch *inputFileFormat {
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

func convertImage(in_path string, out_path string) (err error) {
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

func ConvertImage() error {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return errors.New("error: invalid argument")
	}
	if err := validateFlag(); err != nil {
		return err
	}
	inputFileExt := "." + *inputFileFormat
	outputFileExt := "." + *outputFileFormat
	for _, dir := range args {
		filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", trimError(err))
				return err
			}
			if info.IsDir() || strings.HasSuffix(path, outputFileExt) {
				return nil
			}
			if !strings.HasSuffix(path, inputFileExt) {
				fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
				return nil
			}
			in_path := path
			out_path := replaceSuffix(path, inputFileExt, outputFileExt)
			if err := convertImage(in_path, out_path); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", trimError(err))
				return nil
			}
			return nil
		})
	}
	return nil
}