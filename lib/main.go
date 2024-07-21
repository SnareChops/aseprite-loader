package lib

import (
	"errors"
	"image"

	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/output"
	"github.com/SnareChops/aseprite-loader/transform"
)

type FrameImage = output.FrameImage
type Frame = output.Frame
type Tag = output.Tag
type Layer = output.Layer
type BlendMode = internal.BlendMode

const (
	BlendModeNormal     = internal.BlendModeNormal
	BlendModeMultiply   = internal.BlendModeMultiply
	BlendModeScreen     = internal.BlendModeScreen
	BlendModeOverlay    = internal.BlendModeOverlay
	BlendModeDarken     = internal.BlendModeDarken
	BlendModeLighten    = internal.BlendModeLighten
	BlendModeColorDodge = internal.BlendModeColorDodge
	BlendModeColorBurn  = internal.BlendModeColorBurn
	BlendModeHardLight  = internal.BlendModeHardLight
	BlendModeSoftLight  = internal.BlendModeSoftLight
	BlendModeDifference = internal.BlendModeDifference
	BlendModeExclusion  = internal.BlendModeExclusion
	BlendModeHue        = internal.BlendModeHue
	BlendModeSaturation = internal.BlendModeSaturation
	BlendModeColor      = internal.BlendModeColor
	BlendModeLuminosity = internal.BlendModeLuminosity
	BlendModeAddition   = internal.BlendModeAddition
	BlendModeSubtract   = internal.BlendModeSubtract
	BlendModeDivide     = internal.BlendModeDivide
)

func LoadFrames(path string) ([]FrameImage, []Tag, error) {
	file, err := transform.File(path, "")
	if err != nil {
		return nil, nil, err
	}
	return output.Frames(file)
}

func LoadSplitFramesAndLayers(path string) (frames []Frame, tags []Tag, err error) {
	var file internal.File
	file, err = transform.File(path, "")
	if err != nil {
		return
	}
	frames, tags = output.SplitFramesAndLayers(file)
	return
}

func Smash(layers []Layer) (result image.Image) {
	for _, layer := range layers {
		if !layer.IsVisible {
			continue
		}
		if result == nil {
			result = layer.Image
			continue
		}
		result = output.Blend(result, layer.Image, layer.BlendMode)
	}
	return
}

type Sliceable interface {
	SubImage(r image.Rectangle) image.Image
}

func Slice(img image.Image, gridWidth, gridHeight int) (result []image.Image, err error) {
	if gridWidth == 0 || gridHeight == 0 {
		return nil, errors.New("gridWidth and gridHeight must be greater than 0")
	}
	if sliceable, ok := img.(Sliceable); ok {
		for y := 0; y < img.Bounds().Dy(); y += gridHeight {
			for x := 0; x < img.Bounds().Dx(); x += gridWidth {
				result = append(result, sliceable.SubImage(image.Rect(x, y, x+gridWidth, y+gridHeight)))
			}
		}
		return result, nil
	}
	return nil, errors.New("image does not support slicing")
}

func SmashAndSlice(layers []Layer, gridWidth, gridHeight int) (result []image.Image, err error) {
	return Slice(Smash(layers), gridWidth, gridHeight)
}
