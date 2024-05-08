package transform

import (
	"io"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func transformFrame(in io.Reader, out io.Writer, colorDepth uint16) (frame internal.PreProcessedFrame, err error) {
	trace.Log("transformFrame")
	var header ase.FrameHeader
	header, err = transform[ase.FrameHeader](in, out)
	if err != nil {
		return
	}
	frame.Duration = header.FrameDuration
	count := header.NewNumChunks
	if count == 0 {
		count = uint32(header.OldNumChunks)
	}
	for range count {
		var chunk any
		chunk, err = transformChunk(in, out, colorDepth)
		if err != nil {
			return
		}
		frame.Chunks = append(frame.Chunks, chunk)
	}
	return
}
