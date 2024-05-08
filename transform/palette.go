package transform

import (
	"image/color"
	"io"

	"github.com/SnareChops/aseprite/ase"
	"github.com/SnareChops/aseprite/internal"
	"github.com/SnareChops/aseprite/trace"
)

func transformPalette(in io.Reader, out io.Writer) (data internal.Palette, err error) {
	trace.Log("transformPalette")
	paletteChunk, err := transform[ase.PaletteChunk](in, out)
	if err != nil {
		return
	}
	data.FirstIndex = paletteChunk.FirstIndex
	data.LastIndex = paletteChunk.LastIndex
	for range paletteChunk.Size {
		var palette ase.PaletteEntry
		palette, err = transform[ase.PaletteEntry](in, out)
		if err != nil {
			return
		}
		var name ase.String
		if palette.Flags == 1 {
			name, err = transformString(in, out)
			if err != nil {
				return
			}
		}
		data.Entries = append(data.Entries, internal.PaletteEntry{
			Name:  name.Value,
			Flags: palette.Flags,
			Color: color.NRGBA{palette.Red, palette.Green, palette.Blue, palette.Alpha},
		})
	}
	return
}
