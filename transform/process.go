package transform

import (
	"fmt"
	"image/color"
	"slices"

	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func processFirstFrame(pre internal.PreProcessedFrame) (frame internal.Frame, file internal.FileChunks, err error) {
	trace.Log("processFirstFrame")
	var oldPalette *internal.OldPalette
	var newPalette *internal.Palette
	var prev any
	var userDataIndex int
	for _, chunk := range pre.Chunks {
		if userData, ok := chunk.(internal.UserData); ok {
			processUserData(userData, &prev, userDataIndex)
			userDataIndex++
			continue
		} else {
			userDataIndex = 0
		}
		switch chunk := chunk.(type) {
		case internal.OldPalette:
			oldPalette = &chunk
			prev = chunk
		case internal.Layer:
			file.Layers = append(file.Layers, chunk)
			prev = chunk
		case internal.Cel:
			for len(frame.Cels) <= int(chunk.LayerIndex) {
				frame.Cels = append(frame.Cels, internal.Cel{})
			}
			frame.Cels[chunk.LayerIndex] = chunk
			prev = chunk
		case internal.CelExtra:
			if cel, ok := prev.(internal.Cel); ok {
				cel.CelExtra = chunk
			}
			prev = chunk
		case internal.ColorProfile:
			file.ColorProfile = chunk
			prev = chunk
		case internal.ExternalFiles:
			file.ExternalFiles = chunk
			prev = chunk
		case internal.Tag:
			file.Tags = append(file.Tags, chunk)
			prev = chunk
		case internal.Palette:
			newPalette = &chunk
			prev = chunk
		case internal.Slice:
			file.Slices = append(file.Slices, chunk)
			prev = chunk
		case internal.Tileset:
			file.Tilesets = append(file.Tilesets, chunk)
			prev = chunk
		default:
		}
	}
	if newPalette != nil {
		file.Palette = processNewPalette(*newPalette)
	} else {
		file.Palette = processOldPalette(*oldPalette)
	}
	slices.SortFunc(file.Tilesets, func(a, b internal.Tileset) int {
		return int(b.ID - a.ID)
	})
	return
}

func processFrame(pre internal.PreProcessedFrame) (frame internal.Frame, err error) {
	trace.Log("processFrame")
	var prev any
	var userDataIndex int
	for _, chunk := range pre.Chunks {
		if userData, ok := chunk.(internal.UserData); ok {
			processUserData(userData, &prev, userDataIndex)
			userDataIndex++
			continue
		} else {
			userDataIndex = 0
		}
		switch chunk := chunk.(type) {
		case internal.Cel:
			for len(frame.Cels) <= int(chunk.LayerIndex) {
				frame.Cels = append(frame.Cels, internal.Cel{})
			}
			frame.Cels[chunk.LayerIndex] = chunk
			prev = chunk
		case internal.CelExtra:
			if cel, ok := prev.(internal.Cel); ok {
				cel.CelExtra = chunk
			}
			prev = chunk
		default:
			panic("Unexpected chunk type" + fmt.Sprint(chunk))
		}
	}
	return
}

func processUserData(userData internal.UserData, prev *any, index int) (err error) {
	trace.Log("processUserData")
	switch chunk := (*prev).(type) {
	case internal.OldPalette:
		chunk.UserData = userData
	case internal.Layer:
		chunk.UserData = userData
	case internal.Cel:
		chunk.UserData = userData
	case internal.CelExtra:
		chunk.UserData = userData
	case internal.ColorProfile:
		chunk.UserData = userData
	case internal.ExternalFiles:
		chunk.UserData = userData
	case internal.Mask:
		chunk.UserData = userData
	case internal.Tags:
		chunk[index].UserData = userData
	case internal.Palette:
		chunk.UserData = userData
	case internal.Slice:
		chunk.UserData = userData
	case internal.Tileset:
		if index == 0 {
			chunk.UserData = userData
		} else {
			chunk.Tiles[index-1].UserData = userData
		}
	default:
		err = fmt.Errorf("unexpected user data chunk: %T", chunk)
	}
	return
}

func processOldPalette(chunk internal.OldPalette) (palette []color.Color) {
	trace.Log("processOldPalette")
	for _, packet := range chunk.Packets {
		palette = append(palette, packet.Colors...)
	}
	return
}

func processNewPalette(chunk internal.Palette) (palette []color.Color) {
	trace.Log("processNewPalette")
	for _, entry := range chunk.Entries {
		palette = append(palette, entry.Color)
	}
	return
}
