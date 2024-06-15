package output

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"

	"github.com/SnareChops/aseprite-loader/internal"
)

func Gif(file internal.File, path string) error {
	g, err := createFileAnim(file)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return gif.EncodeAll(f, g)
}

type ImageFrame struct {
	Duration int
	Image    image.Image
}

func createFileAnim(file internal.File) (result *gif.GIF, err error) {
	result = &gif.GIF{}
	for _, frame := range file.Frames {
		var im image.Image
		im, err = CreateFrameImage(file, frame)
		if err != nil {
			return
		}
		pimg := image.NewPaletted(im.Bounds(), palette.Plan9)
		options := gif.Options{
			NumColors: 256,
			Drawer:    draw.FloydSteinberg,
		}
		if options.Quantizer != nil {
			pimg.Palette = options.Quantizer.Quantize(make(color.Palette, 0, 256), im)
		}
		options.Drawer.Draw(pimg, im.Bounds(), im, image.Point{})
		result.Image = append(result.Image, pimg)
		result.Delay = append(result.Delay, int(frame.Duration)/10)
	}
	return
}
