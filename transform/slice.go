package transform

import (
	"io"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func transformSlice(in io.Reader, out io.Writer) (data internal.Slice, err error) {
	trace.Log("transformSlice")
	var sliceChunk ase.SliceChunk
	sliceChunk, err = transform[ase.SliceChunk](in, out)
	if err != nil {
		return
	}
	data.Flags = sliceChunk.Flags
	var name ase.String
	name, err = transformString(in, out)
	if err != nil {
		return
	}
	data.Name = name.Value
	for range sliceChunk.Keys {
		var key internal.SliceKey
		key, err = transformSliceKey(in, out, sliceChunk.Flags)
		data.Keys = append(data.Keys, key)
	}
	return
}

func transformSliceKey(in io.Reader, out io.Writer, flags uint32) (data internal.SliceKey, err error) {
	trace.Log("transformSliceKey")
	var key ase.SliceKey
	key, err = transform[ase.SliceKey](in, out)
	if err != nil {
		return
	}
	data.Frame = key.Frame
	data.X = key.X
	data.Y = key.Y
	data.Width = key.Width
	data.Height = key.Height
	if flags&0x1 == 0x1 {
		var patches ase.Slice9Patches
		patches, err = transform[ase.Slice9Patches](in, out)
		if err != nil {
			return
		}
		data.Patches = internal.Slice9Patches{
			X:      patches.X,
			Y:      patches.Y,
			Width:  patches.Width,
			Height: patches.Height,
		}
	}
	if flags&0x2 == 0x2 {
		var pivot ase.SlicePivot
		pivot, err = transform[ase.SlicePivot](in, out)
		if err != nil {
			return
		}
		data.Pivot = internal.SlicePivot{
			X: pivot.X,
			Y: pivot.Y,
		}
	}
	return
}
