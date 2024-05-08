package transform

import (
	"image/color"
	"io"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func transformOldPalette(in io.Reader, out io.Writer) (palette internal.OldPalette, err error) {
	trace.Log("transformOldPalette")
	var header ase.OldPaletteHeader
	header, err = transform[ase.OldPaletteHeader](in, out)
	if err != nil {
		return
	}
	for range header.NumberOfPackets {
		var packet ase.OldPalettePacket
		packet, err = transform[ase.OldPalettePacket](in, out)
		if err != nil {
			return
		}
		pack := internal.OldPacket{
			Colors: []color.Color{},
		}
		for range packet.NumberOfColors {
			var clr ase.OldPacketColor
			clr, err = transform[ase.OldPacketColor](in, out)
			if err != nil {
				return
			}
			pack.Colors = append(pack.Colors, color.NRGBA{clr.Red, clr.Green, clr.Blue, 0xff})
		}
		palette.Packets = append(palette.Packets, pack)
	}
	return
}
