package output

import (
	"image"

	"github.com/SnareChops/aseprite/internal"
)

type Frame struct {
	Width    int
	Height   int
	Duration int
	Layers   []Layer
}

type Layer struct {
	Name      string
	BlendMode internal.BlendMode
	Opacity   byte
	Image     image.Image
}

type FrameImage struct {
	Duration int
	Image    image.Image
}

func SplitFramesAndLayers(file internal.File) (frames []Frame) {
	for _, frame := range file.Frames {
		frames = append(frames, Frame{
			Duration: int(frame.Duration),
			Layers:   LayersForFrame(file, frame),
		})
	}
	return
}

func LayersForFrame(file internal.File, frame internal.Frame) (layers []Layer) {
	for _, layer := range frame.Layers {
		layers = append(layers, Layer{
			Name:      layer.Name,
			BlendMode: layer.BlendMode,
			Opacity:   layer.Opacity,
			Image:     CreateLayerImage(file, layer),
		})
	}
	return
}

func Frames(file internal.File) (frames []FrameImage, err error) {
	for _, frame := range file.Frames {
		var image image.Image
		image, err = CreateFrameImage(file, frame)
		if err != nil {
			return
		}
		frames = append(frames, FrameImage{
			Duration: int(frame.Duration),
			Image:    image,
		})
	}
	return
}
