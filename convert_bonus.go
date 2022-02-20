package main

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
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
