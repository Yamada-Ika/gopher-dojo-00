package main

import (
	"fmt"

	"example.com/ex00/imgconv"
)

func main() {
	if err := imgconv.JpgToPng(); err != nil {
		fmt.Println(err.Error())
	}
}
