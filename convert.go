package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

type Image image.Image

func writeImage(file io.Writer, img Image) (err error) {
	err = png.Encode(file, img)
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
