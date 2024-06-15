package main

import (
	"os"

	"github.com/SnareChops/aseprite-loader/output"
	"github.com/SnareChops/aseprite-loader/transform"
)

func main() {
	if len(os.Args) < 2 {
		println("Missing input argument")
	}
	if len(os.Args) < 3 {
		println("Missing output argument")
	}
	file, err := transform.File(input, debug)
	if err != nil {
		panic(err)
	}

	if png != "" {
		err = output.Png(file, png)
		if err != nil {
			panic(err)
		}
	}
	if gif != "" {
		err = output.Gif(file, gif)
		if err != nil {
			panic(err)
		}
	}
}
