package transform

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/SnareChops/aseprite/ase"
	"github.com/SnareChops/aseprite/internal"
	"github.com/SnareChops/aseprite/trace"
)

func transformCel(in io.Reader, out io.Writer, chunkSize uint32, colorDepth uint16) (data internal.Cel, err error) {
	trace.Log("transformCel")
	var header ase.CelChunk
	header, err = transform[ase.CelChunk](in, out)
	if err != nil {
		return
	}
	data.LayerIndex = header.Index
	data.X = int(header.XPosition)
	data.Y = int(header.YPosition)
	data.Opacity = header.Opacity
	data.Type = internal.CelType(header.Type)
	data.ZIndex = header.ZIndex
	remainingSize := chunkSize - 16
	switch data.Type {
	case 0:
		data.Image, err = transformRawImage(in, out, colorDepth)
	case 1:
		data.Link, err = transform[uint16](in, out)
	case 2:
		data.Image, err = transformCompressedImage(in, out, remainingSize)
	case 3:
		data.Tilemap, err = transformCompressedTilemap(in, out, remainingSize)
	}
	return
}

func transformRawImage(in io.Reader, out io.Writer, colorDepth uint16) (data internal.CelImage, err error) {
	trace.Log("transformRawImage")
	var header ase.RawImageCel
	header, err = transform[ase.RawImageCel](in, out)
	if err != nil {
		return
	}
	data.Width = int(header.Width)
	data.Height = int(header.Height)
	data.Bytes, err = transformBytes(in, out, header.Width*header.Height*(colorDepth/8))
	return
}

func transformCompressedImage(in io.Reader, out io.Writer, size uint32) (data internal.CelImage, err error) {
	trace.Log("transformCompressedImage")
	var header ase.CompressedImageCel
	header, err = transform[ase.CompressedImageCel](in, out)
	if err != nil {
		return
	}
	data.Width = int(header.Width)
	data.Height = int(header.Height)
	var compressed []byte
	compressed, err = transformBytes(in, out, size-4)
	if err != nil {
		return
	}
	data.Bytes, err = decompress(compressed)
	return
}

func transformCompressedTilemap(in io.Reader, out io.Writer, size uint32) (data internal.CelTilemap, err error) {
	trace.Log("transformCompressedTilemap")
	var header ase.CompressedTilemapCel
	header, err = transform[ase.CompressedTilemapCel](in, out)
	if err != nil {
		return
	}
	data.Width = header.Width
	data.Height = header.Height
	data.BitsPerTile = header.BitsPerTile
	data.TileIDMask = header.TileIDMask
	data.XFlipMask = header.XFlipMask
	data.YFlipMask = header.YFlipMask
	data.DFlipMask = header.DFlipMask
	var compressed []byte
	compressed, err = transformBytes(in, out, size-32)
	if err != nil {
		return
	}
	raw, err := decompress(compressed)
	if err != nil {
		return
	}
	for i := 0; i < len(raw); i += 4 {
		var tile uint32
		err = binary.Read(bytes.NewReader(raw[i:i+4]), binary.LittleEndian, &tile)
		data.Tiles = append(data.Tiles, tile)
	}
	return
}
