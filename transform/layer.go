package transform

import (
	"io"

	"github.com/SnareChops/aseprite/ase"
	"github.com/SnareChops/aseprite/internal"
	"github.com/SnareChops/aseprite/trace"
)

func transformLayer(in io.Reader, out io.Writer) (layer internal.Layer, err error) {
	trace.Log("transformLayer")
	var header ase.LayerChunk
	header, err = transform[ase.LayerChunk](in, out)
	if err != nil {
		return
	}
	layer.Flags = internal.LayerFlag(header.Flags)
	layer.Type = header.Type
	layer.DefaultWidth = header.DefaultWidth
	layer.DefaultHeight = header.DefaultHeight
	layer.BlendMode = internal.BlendMode(header.BlendMode)
	layer.Opacity = header.Opacity
	var name ase.String
	name, err = transformString(in, out)
	if err != nil {
		return
	}
	layer.Name = name.Value
	if layer.Type == 2 {
		var id uint32
		id, err = transform[uint32](in, out)
		if err != nil {
			return
		}
		layer.TilesetID = id
	}
	return
}
