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

func LoadFrames(path string) ([]FrameImage, error) {
	file, err := transform.File(path, "")
	if err != nil {
		return nil, err
	}
	return output.Frames(file)
}

func LoadSplitFramesAndLayers(path string) ([]Frame, error) {
	file, err := transform.File(path, "")
	if err != nil {
		return nil, err
	}
	return output.SplitFramesAndLayers(file), nil
}

func Smash(layers []Layer) image.Image {
	var im image.Image = image.NewNRGBA(image.Rect(0, 0, 1, 1))
	for _, layer := range layers {
		if !layer.IsVisible {
			continue
		}
		im = output.Blend(im, layer.Image, layer.BlendMode)
	}
	return im
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
