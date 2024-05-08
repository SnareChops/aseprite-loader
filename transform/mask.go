package transform

import (
	"io"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func transformMask(in io.Reader, out io.Writer) (data internal.Mask, err error) {
	trace.Log("transformMask")
	var mask ase.MaskChunk
	mask, err = transform[ase.MaskChunk](in, out)
	if err != nil {
		return
	}
	data.X = mask.X
	data.Y = mask.Y
	data.Width = mask.Width
	data.Height = mask.Height
	var name ase.String
	name, err = transformString(in, out)
	if err != nil {
		return
	}
	data.Name = name.Value
	size := mask.Width * ((mask.Width + 7) / 8)
	data.Data, err = transformBytes(in, out, uint32(size))
	return
}
