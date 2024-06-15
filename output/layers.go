package output

import (
	"image"

	"github.com/SnareChops/aseprite-loader/internal"
)

type Frame struct {
	Width      int
	Height     int
	Duration   int
	GridWidth  int
	GridHeight int
	Layers     []Layer
}

type Layer struct {
	Name      string
	BlendMode internal.BlendMode
	Opacity   byte
	IsVisible bool
	Image     image.Image
}

type FrameImage struct {
	Duration   int
	Image      image.Image
	GridWidth  int
	GridHeight int
}

func SplitFramesAndLayers(file internal.File) (frames []Frame) {
	for _, frame := range file.Frames {
		frames = append(frames, Frame{
			Width:      file.Width,
			Height:     file.Height,
			GridWidth:  int(file.GridWidth),
			GridHeight: int(file.GridHeight),
			Duration:   int(frame.Duration),
			Layers:     LayersForFrame(file, frame),
		})
	}
	return
}

func LayersForFrame(file internal.File, frame internal.Frame) (layers []Layer) {
	for i, layer := range file.Layers {
		layers = append(layers, Layer{
			Name:      layer.Name,
			BlendMode: layer.BlendMode,
			Opacity:   layer.Opacity,
			IsVisible: layer.Flags&internal.LayerFlagVisible != 0,
			Image:     CreateCelImage(file, frame.Cels[i], i),
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
			Duration:   int(frame.Duration),
			Image:      image,
			GridWidth:  int(file.GridWidth),
			GridHeight: int(file.GridHeight),
		})
	}
	return
}
