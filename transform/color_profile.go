package transform

import (
	"io"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func transformColorProfile(in io.Reader, out io.Writer) (data internal.ColorProfile, err error) {
	trace.Log("transformColorProfile")
	colorProfileChunk, err := transform[ase.ColorProfileChunk](in, out)
	if err != nil {
		return
	}
	data.Type = colorProfileChunk.Type
	data.Flags = colorProfileChunk.Flags
	data.Gamma = colorProfileChunk.Gamma
	if colorProfileChunk.Flags == 2 {
		var length uint32
		length, err = transform[uint32](in, out)
		if err != nil {
			return
		}
		var icc []byte
		icc, err = transformBytes(in, out, length)
		if err != nil {
			return
		}
		data.ICC = icc
	}
	return
}
