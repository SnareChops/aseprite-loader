package lib

import (
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

func SmashAndSlice(layers []Layer, gridWidth, gridHeight int) (result []image.Image) {
	im := Smash(layers)
	for y := 0; y < im.Bounds().Dy(); y += gridHeight {
		for x := 0; x < im.Bounds().Dx(); x += gridWidth {
			result = append(result, im.(*image.NRGBA).SubImage(image.Rect(x, y, x+gridWidth, y+gridHeight)))
		}
	}
	return
}
