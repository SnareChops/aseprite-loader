package lib

import (
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
