package transform

import (
	"image"
	"image/color"
	"io"

	"github.com/SnareChops/aseprite/ase"
	"github.com/SnareChops/aseprite/internal"
	"github.com/SnareChops/aseprite/trace"
)

func transformTileset(in io.Reader, out io.Writer) (data internal.Tileset, err error) {
	trace.Log("transformTileset")
	tileset, err := transform[ase.TilesetChunk](in, out)
	if err != nil {
		return
	}
	var name ase.String
	name, err = transformString(in, out)
	if err != nil {
		return
	}
	data.ID = tileset.ID
	data.Name = name.Value
	data.Flags = tileset.Flags
	data.Width = tileset.Width
	data.Height = tileset.Height
	data.Index = tileset.Index
	if tileset.Flags&0x1 == 0x1 {
		var external ase.TilesetExternalData
		external, err = transform[ase.TilesetExternalData](in, out)
		if err != nil {
			return
		}
		data.ExternalID = external.ID
		data.ExternalTilesetID = external.TilesetID
	}
	if tileset.Flags&0x2 == 0x2 {
		var tiles []image.Image
		tiles, err = transformTilesetCompressedData(in, out, int(tileset.Width), int(tileset.Height))
		if err != nil {
			return
		}
		for _, tile := range tiles {
			data.Tiles = append(data.Tiles, internal.Tile{Image: tile})
		}
	}
	return
}

func transformTilesetCompressedData(in io.Reader, out io.Writer, width int, height int) (data []image.Image, err error) {
	trace.Log("transformTilesetCompressedData")
	var length uint32
	length, err = transform[uint32](in, out)
	if err != nil {
		return
	}
	var compressed []byte
	compressed, err = transformBytes(in, out, length)
	if err != nil {
		return
	}
	var pix []byte
	pix, err = decompress(compressed)
	if err != nil {
		return
	}
	for t := 0; t < len(pix); t += width * height * 4 {
		img := image.NewNRGBA(image.Rect(0, 0, width, height))
		for i := 0; i < width*height*4; i += 4 {
			c := color.NRGBA{pix[t+i], pix[t+i+1], pix[t+i+2], pix[t+i+3]}
			img.Set((i/4)%width, (i/4)/width, c)
		}
		data = append(data, img)
	}
	return
}
