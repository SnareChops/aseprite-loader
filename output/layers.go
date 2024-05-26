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
	Duration int
	Image    image.Image
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
	for _, layer := range frame.Layers {
		layers = append(layers, Layer{
			Name:      layer.Name,
			BlendMode: layer.BlendMode,
			Opacity:   layer.Opacity,
			IsVisible: layer.Flags&internal.LayerFlagVisible != 0,
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

func Smash(layers []Layer) image.Image {
	var im image.Image = image.NewNRGBA(image.Rect(0, 0, 1, 1))
	for _, layer := range layers {
		if !layer.IsVisible {
			continue
		}
		im = Blend(im, layer.Image, layer.BlendMode)
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
