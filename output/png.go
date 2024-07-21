package output

import (
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
	for i, layer := range SortLayers(file.Layers, frame.Cels) {
		if layer.Flags&internal.LayerFlagVisible == 0 {
			continue
		}
		if len(frame.Cels) > i {
			im = Blend(im, CreateCelImage(file, frame.Cels[i], i), layer.BlendMode)
		}
	}
	return im, nil
}

func Blend(dest, src image.Image, mode internal.BlendMode) image.Image {
	switch mode {
	case internal.BlendModeNormal:
		return blend.Normal(dest, src)
	case internal.BlendModeMultiply:
		return blend.Multiply(dest, src)
	case internal.BlendModeScreen:
		return blend.Screen(dest, src)
	case internal.BlendModeOverlay:
		return blend.Overlay(dest, src)
	case internal.BlendModeDarken:
		return blend.Darken(dest, src)
	case internal.BlendModeLighten:
		return blend.Lighten(dest, src)
	case internal.BlendModeColorDodge:
		return blend.LinearDodge(dest, src)
	case internal.BlendModeColorBurn:
		return blend.LinearBurn(dest, src)
	case internal.BlendModeHardLight:
		return blend.HardLight(dest, src)
	case internal.BlendModeSoftLight:
		return blend.SoftLight(dest, src)
	case internal.BlendModeDifference:
		return blend.Difference(dest, src)
	case internal.BlendModeExclusion:
		return blend.Exclusion(dest, src)
	case internal.BlendModeHue:
		return blend.Hue(dest, src)
	case internal.BlendModeSaturation:
		return blend.Saturation(dest, src)
	case internal.BlendModeColor:
		return blend.Color(dest, src)
	case internal.BlendModeLuminosity:
		return blend.Luminosity(dest, src)
	case internal.BlendModeAddition:
		return blend.Addition(dest, src)
	case internal.BlendModeSubtract:
		return blend.Subtraction(dest, src)
	case internal.BlendModeDivide:
		return Divide(dest, src)
	default:
		return nil
	}
}

func CreateCelImage(file internal.File, cel internal.Cel, index int) image.Image {
	trace.Log("createLayerImage")
	switch cel.Type {
	case internal.CelTypeRawImage, internal.CelTypeCompressedImage:
		return createImageFromRaw(file, cel, index)
	case internal.CelTypeLinkedCel:
		return createImageFromLinked(file, cel, index)
	case internal.CelTypeCompressedTilemap:
		return createImageFromTilemap(file, cel, index)
	default:
		panic("unknown cel type")
	}
}

func createImageFromRaw(file internal.File, cel internal.Cel, index int) image.Image {
	trace.Log("createImageFromRaw")
	switch file.ColorDepth {
	case 8:
		return createPaletteImage(file, cel)
	case 16:
		return CreateGrayscaleImage(file, cel, index)
	case 32:
		return CreateRGBAImage(file, cel, index)
	default:
		panic("unknown color depth")
	}
}

func createPaletteImage(file internal.File, cel internal.Cel) image.Image {
	trace.Log("createPaletteImage")
	result := image.NewNRGBA(image.Rect(0, 0, file.Width, file.Height))
	celImage := cel.Image
	for i := range celImage.Width * celImage.Height {
		result.Set(i%file.Width+cel.X, i/file.Width+cel.Y, file.Palette[celImage.Bytes[i]])
	}
	return result
}

func CreateGrayscaleImage(file internal.File, cel internal.Cel, index int) image.Image {
	trace.Log("createGrayscaleImage")
	result := image.NewNRGBA(image.Rect(0, 0, file.Width, file.Height))
	celImage := cel.Image
	for i := 0; i < celImage.Width*celImage.Height*2; i += 2 {
		c := color.NRGBA{celImage.Bytes[i], celImage.Bytes[i], celImage.Bytes[i], celImage.Bytes[i+1]}
		result.Set((i/2)%file.Width+cel.X, (i/2)/file.Width+cel.Y, c)
	}
	return result
}

func CreateRGBAImage(file internal.File, cel internal.Cel, index int) image.Image {
	trace.Log("createRGBAImage")
	result := image.NewNRGBA(image.Rect(0, 0, file.Width, file.Height))
	celImage := cel.Image
	for i := 0; i < celImage.Width*celImage.Height*4; i += 4 {
		c := color.NRGBA{celImage.Bytes[i], celImage.Bytes[i+1], celImage.Bytes[i+2], celImage.Bytes[i+3]}
		x := (i/4)%celImage.Width + cel.X
		y := (i/4)/celImage.Width + cel.Y
		if x < 0 || y < 0 || x >= file.Width || y >= file.Height {
			continue
		}
		result.Set(x, y, c)
	}
	return result
}

func createImageFromLinked(file internal.File, cel internal.Cel, index int) image.Image {
	trace.Log("createImageFromLinked")
	return CreateCelImage(file, file.Frames[cel.Link].Cels[cel.LayerIndex], index)
}

func createImageFromTilemap(file internal.File, cel internal.Cel, index int) image.Image {
	trace.Log("createImageFromTilemap")
	result := image.NewNRGBA(image.Rect(0, 0, file.Width, file.Height))
	tilemap := cel.Tilemap
	tileset := file.Tilesets[file.Layers[index].TilesetID]
	for i, t := range cel.Tilemap.Tiles {
		tile := tileset.Tiles[t&tilemap.TileIDMask]
		for y := range tile.Image.Bounds().Dy() {
			for x := range tile.Image.Bounds().Dx() {
				result.Set(i%file.Width+cel.X+x, i/file.Width+cel.Y+y, tile.Image.At(x, y))
			}
		}
	}
	return result
}

func SortLayers(layers []internal.Layer, cels []internal.Cel) []internal.Layer {
	type indexed struct {
		index int
		layer internal.Layer
		cel   internal.Cel
	}
	pre := make([]indexed, len(layers))
	for i, layer := range layers {
		if len(cels) > i {
			pre[i] = indexed{i, layer, cels[i]}
		} else {
			pre[i] = indexed{i, layer, internal.Cel{}}
		}
	}
	slices.SortStableFunc(pre, func(a, b indexed) int {
		av := a.index + int(a.cel.ZIndex)
		bv := b.index + int(b.cel.ZIndex)
		return av - bv
	})
	result := make([]internal.Layer, len(pre))
	for i, p := range pre {
		result[i] = p.layer
	}
	return result
}
