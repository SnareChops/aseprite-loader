package output

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
	"slices"

	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
	"hawx.me/code/img/blend"
)

func Png(file internal.File, path string) error {
	image, err := createFileImage(file)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, image)
}

func createFileImage(file internal.File) (image.Image, error) {
	im := image.NewNRGBA(image.Rect(0, 0, file.Width*len(file.Frames), file.Height))
	for i, frame := range file.Frames {
		frameImage, err := CreateFrameImage(file, frame)
		if err != nil {
			return nil, err
		}
		for y := range frameImage.Bounds().Dy() {
			for x := range frameImage.Bounds().Dx() {
				im.Set(x+i*file.Width, y, frameImage.At(x, y))
			}
		}
	}
	return im, nil
}

func CreateFrameImage(file internal.File, frame internal.Frame) (image.Image, error) {
	var im image.Image = image.NewNRGBA(image.Rect(0, 0, file.Width, file.Height))
	for _, layer := range SortLayers(frame.Layers) {
		if layer.Flags&internal.LayerFlagVisible == 0 {
			continue
		}
		// TODO switch on blend mode
		switch layer.BlendMode {
		case internal.BlendModeNormal:
			im = blend.Normal(im, CreateLayerImage(file, layer))
		case internal.BlendModeMultiply:
			im = blend.Multiply(im, CreateLayerImage(file, layer))
		case internal.BlendModeScreen:
			im = blend.Screen(im, CreateLayerImage(file, layer))
		case internal.BlendModeOverlay:
			im = blend.Overlay(im, CreateLayerImage(file, layer))
		case internal.BlendModeDarken:
			im = blend.Darken(im, CreateLayerImage(file, layer))
		case internal.BlendModeLighten:
			im = blend.Lighten(im, CreateLayerImage(file, layer))
		case internal.BlendModeColorDodge:
			im = blend.LinearDodge(im, CreateLayerImage(file, layer))
		case internal.BlendModeColorBurn:
			im = blend.LinearBurn(im, CreateLayerImage(file, layer))
		case internal.BlendModeHardLight:
			im = blend.HardLight(im, CreateLayerImage(file, layer))
		case internal.BlendModeSoftLight:
			im = blend.SoftLight(im, CreateLayerImage(file, layer))
		case internal.BlendModeDifference:
			im = blend.Difference(im, CreateLayerImage(file, layer))
		case internal.BlendModeExclusion:
			im = blend.Exclusion(im, CreateLayerImage(file, layer))
		case internal.BlendModeHue:
			im = blend.Hue(im, CreateLayerImage(file, layer))
		case internal.BlendModeSaturation:
			im = blend.Saturation(im, CreateLayerImage(file, layer))
		case internal.BlendModeColor:
			im = blend.Color(im, CreateLayerImage(file, layer))
		case internal.BlendModeLuminosity:
			im = blend.Luminosity(im, CreateLayerImage(file, layer))
		case internal.BlendModeAddition:
			im = blend.Addition(im, CreateLayerImage(file, layer))
		case internal.BlendModeSubtract:
			im = blend.Subtraction(im, CreateLayerImage(file, layer))
		case internal.BlendModeDivide:
			im = Divide(im, CreateLayerImage(file, layer))
		default:
			return nil, errors.New("unknown blend mode")
		}
	}
	return im, nil
}

func CreateLayerImage(file internal.File, layer internal.Layer) image.Image {
	trace.Log("createLayerImage")
	switch layer.Cel.Type {
	case internal.CelTypeRawImage, internal.CelTypeCompressedImage:
		return createImageFromRaw(file, layer)
	case internal.CelTypeLinkedCel:
		return createImageFromLinked(file, layer)
	case internal.CelTypeCompressedTilemap:
		return createImageFromTilemap(file, layer)
	default:
		panic("unknown cel type")
	}
}

func createImageFromRaw(file internal.File, layer internal.Layer) image.Image {
	trace.Log("createImageFromRaw")
	switch file.ColorDepth {
	case 8:
		return createPaletteImage(file, layer)
	case 16:
		return CreateGrayscaleImage(file, layer)
	case 32:
		return CreateRGBAImage(file, layer)
	default:
		panic("unknown color depth")
	}
}

func createPaletteImage(file internal.File, layer internal.Layer) image.Image {
	trace.Log("createPaletteImage")
	result := image.NewNRGBA(image.Rect(0, 0, file.Width, file.Height))
	celImage := layer.Cel.Image
	for i := range celImage.Width * celImage.Height {
		result.Set(i%file.Width+layer.Cel.X, i/file.Width+layer.Cel.Y, file.Palette[celImage.Bytes[i]])
	}
	return result
}

func CreateGrayscaleImage(file internal.File, layer internal.Layer) image.Image {
	trace.Log("createGrayscaleImage")
	result := image.NewNRGBA(image.Rect(0, 0, file.Width, file.Height))
	celImage := layer.Cel.Image
	for i := 0; i < celImage.Width*celImage.Height*2; i += 2 {
		c := color.NRGBA{celImage.Bytes[i], celImage.Bytes[i], celImage.Bytes[i], celImage.Bytes[i+1]}
		result.Set((i/2)%file.Width+layer.Cel.X, (i/2)/file.Width+layer.Cel.Y, c)
	}
	return result
}

func CreateRGBAImage(file internal.File, layer internal.Layer) image.Image {
	trace.Log("createRGBAImage")
	result := image.NewNRGBA(image.Rect(0, 0, file.Width, file.Height))
	celImage := layer.Cel.Image
	for i := 0; i < celImage.Width*celImage.Height*4; i += 4 {
		c := color.NRGBA{celImage.Bytes[i], celImage.Bytes[i+1], celImage.Bytes[i+2], celImage.Bytes[i+3]}
		result.Set((i/4)%file.Width+layer.Cel.X, (i/4)/file.Width+layer.Cel.Y, c)
	}
	return result
}

func createImageFromLinked(file internal.File, layer internal.Layer) image.Image {
	trace.Log("createImageFromLinked")
	return file.Frames[layer.Cel.Link].Layers[layer.Cel.LayerIndex].Image
}

func createImageFromTilemap(file internal.File, layer internal.Layer) image.Image {
	trace.Log("createImageFromTilemap")
	result := image.NewNRGBA(image.Rect(0, 0, file.Width, file.Height))
	tilemap := layer.Cel.Tilemap
	tileset := file.Tilesets[layer.TilesetID]
	for i, t := range layer.Cel.Tilemap.Tiles {
		tile := tileset.Tiles[t&tilemap.TileIDMask]
		for y := range tile.Image.Bounds().Dy() {
			for x := range tile.Image.Bounds().Dx() {
				result.Set(i%file.Width+layer.Cel.X+x, i/file.Width+layer.Cel.Y+y, tile.Image.At(x, y))
			}
		}
	}
	return result
}

func SortLayers(layers []internal.Layer) []internal.Layer {
	type indexed struct {
		index int
		layer internal.Layer
	}
	pre := make([]indexed, len(layers))
	for i, layer := range layers {
		pre[i] = indexed{i, layer}
	}
	slices.SortStableFunc(pre, func(a, b indexed) int {
		av := a.index + int(a.layer.Cel.ZIndex)
		bv := b.index + int(b.layer.Cel.ZIndex)
		return av - bv
	})
	result := make([]internal.Layer, len(pre))
	for i, p := range pre {
		result[i] = p.layer
	}
	return result
}
