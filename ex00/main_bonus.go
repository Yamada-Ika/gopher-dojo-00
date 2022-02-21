package main

import (
	"fmt"

	"example.com/ex00/imgconv_bonus"
)

func main() {
	if err := imgconv_bonus.ConvertImage(); err != nil {
		fmt.Println(err.Error())
	}
}
