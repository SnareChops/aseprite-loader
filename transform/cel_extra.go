package transform

import (
	"io"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func transformCelExtra(in io.Reader, out io.Writer) (data internal.CelExtra, err error) {
	trace.Log("transformCelExtra")
	var header ase.CelExtraChunk
	header, err = transform[ase.CelExtraChunk](in, out)
	if err != nil {
		return
	}
	data.Flags = header.Flags
	data.X = header.XPosition
	data.Y = header.YPosition
	data.Width = header.Width
	data.Height = header.Height
	return
}
