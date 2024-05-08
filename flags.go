package main

import (
	"flag"
	"os"

	"github.com/SnareChops/aseprite/trace"
)

var debug string
var png string
var input string
var Trace bool

func init() {
	flag.StringVar(&debug, "debug", "", "Enable debug aseprite parsing output")
	flag.StringVar(&png, "png", "", "Output to PNG file")
	flag.BoolVar(&Trace, "trace", false, "Enable trace output")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		println("Missing input argument")
		flag.Usage()
		os.Exit(1)
	}
	input = args[0]
	trace.Enabled = Trace
}
