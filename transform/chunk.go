package transform

import (
	"io"

	"github.com/SnareChops/aseprite/ase"
	"github.com/SnareChops/aseprite/trace"
)

func transformChunk(in io.Reader, out io.Writer, colorDepth uint16) (data any, err error) {
	trace.Log("transformChunk")
	var header ase.ChunkHeader
	header, err = transform[ase.ChunkHeader](in, out)
	if err != nil {
		return
	}
	switch header.ChunkType {
	case ase.ChunkTypeOldPalette, ase.ChunkTypeAltPalette:
		data, err = transformOldPalette(in, out)
	case ase.ChunkTypeLayer:
		data, err = transformLayer(in, out)
	case ase.ChunkTypeCel:
		data, err = transformCel(in, out, header.ChunkSize-6, colorDepth)
	case ase.ChunkTypeCelExtra:
		data, err = transformCelExtra(in, out)
	case ase.ChunkTypeColorProfile:
		data, err = transformColorProfile(in, out)
	case ase.ChunkTypeExternalFiles:
		data, err = transformExternalFiles(in, out)
	case ase.ChunkTypeMask:
		data, err = transformMask(in, out)
	case ase.ChunkTypeTags:
		data, err = transformTags(in, out)
	case ase.ChunkTypePalette:
		data, err = transformPalette(in, out)
	case ase.ChunkTypeUserData:
		data, err = transformUserData(in, out)
	case ase.ChunkTypeSlice:
		data, err = transformSlice(in, out)
	case ase.ChunkTypeTileset:
		data, err = transformTileset(in, out)
	}
	return
}
