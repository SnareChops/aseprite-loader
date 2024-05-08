package transform

import (
	"io"
	"os"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func File(input string, output string) (file internal.File, err error) {
	var in, out *os.File
	in, err = os.Open(input)
	if err != nil {
		return
	}
	defer in.Close()
	if output != "" {
		out, err = os.Create(output)
	}
	defer out.Close()
	file, err = transformFile(in, out)
	return
}

func transformFile(in io.Reader, out io.Writer) (file internal.File, err error) {
	trace.Log("transformFile", out)
	var header ase.FileHeader
	header, err = transform[ase.FileHeader](in, out)
	if err != nil {
		return
	}
	file.Width = int(header.Width)
	file.Height = int(header.Height)
	file.ColorDepth = header.ColorDepth
	file.PaletteIndex = header.PaletteIndex
	file.NumberOfColors = header.NumberOfColors
	file.PixelWidth = header.PixelWidth
	file.PixelHeight = header.PixelHeight
	file.XGridPosition = header.XGridPosition
	file.YGridPosition = header.YGridPosition
	file.GridWidth = header.GridWidth
	file.GridHeight = header.GridHeight
	for i := range header.Frames {
		var pre internal.PreProcessedFrame
		pre, err = transformFrame(in, out, file.ColorDepth)
		if err != nil {
			return
		}
		var frame internal.Frame
		var fileChunks internal.FileChunks
		frame, fileChunks, err = processFrame(pre)
		if err != nil {
			return
		}
		if i == 0 {
			file.FileChunks = fileChunks
		}
		file.Frames = append(file.Frames, frame)
	}
	return
}
