package lib

import (
	"github.com/SnareChops/aseprite-loader/output"
	"github.com/SnareChops/aseprite-loader/transform"
)

type FrameImage output.FrameImage
type Frame output.Frame
type Layer output.Layer

func LoadFrames(path string) ([]output.FrameImage, error) {
	file, err := transform.File(path, "")
	if err != nil {
		return nil, err
	}
	return output.Frames(file)
}

func LoadSplitFramesAndLayers(path string) ([]output.Frame, error) {
	file, err := transform.File(path, "")
	if err != nil {
		return nil, err
	}
	return output.SplitFramesAndLayers(file), nil
}
