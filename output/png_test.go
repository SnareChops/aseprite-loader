package output_test

import (
	"image/color"
	"testing"

	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/output"
)

func TestCreateRGBAImage(t *testing.T) {
	file := internal.File{
		Width:  2,
		Height: 2,
	}
	cel := internal.Cel{
		Image: internal.CelImage{
			Width:  2,
			Height: 2,
			Bytes: []byte{
				255, 0, 0, 255,
				0, 255, 0, 255,
				0, 0, 255, 255,
				0, 0, 0, 255,
			},
		},
	}
	image := output.CreateRGBAImage(file, cel, 0)
	if image.Bounds().Dx()*image.Bounds().Dy() != 4 {
		t.Fatal("expected 4 pixels")
	}
	if matchesColor(image.At(0, 0), 255, 0, 0, 255) {
		t.Fatal("expected red pixel")
	}
	if matchesColor(image.At(1, 0), 0, 255, 0, 255) {
		t.Fatal("expected green pixel")
	}
	if matchesColor(image.At(0, 1), 0, 0, 255, 255) {
		t.Fatal("expected blue pixel")
	}
	if matchesColor(image.At(1, 1), 0, 0, 0, 255) {
		t.Fatal("expected black pixel")
	}
}

func TestCreateGrayscaleImage(t *testing.T) {
	file := internal.File{
		Width:  2,
		Height: 2,
	}
	cel := internal.Cel{
		Image: internal.CelImage{
			Width:  2,
			Height: 2,
			Bytes: []byte{
				0, 255,
				50, 255,
				100, 255,
				150, 255,
			},
		},
	}
	image := output.CreateGrayscaleImage(file, cel, 0)
	if image.Bounds().Dx()*image.Bounds().Dy() != 4 {
		t.Fatal("expected 4 pixels")
	}
	if matchesColor(image.At(0, 0), 0, 0, 0, 255) {
		t.Fatal("expected black pixel")
	}
	if matchesColor(image.At(1, 0), 50, 50, 50, 255) {
		t.Fatal("expected 50 pixel")
	}
	if matchesColor(image.At(0, 1), 100, 100, 100, 255) {
		t.Fatal("expected 100 pixel")
	}
	if matchesColor(image.At(1, 1), 150, 150, 150, 255) {
		t.Fatal("expected 150 pixel")
	}
}

func TestSortLayers(t *testing.T) {
	layers := []internal.Layer{
		{Name: "0"},
		{Name: "1"},
		{Name: "2"},
		{Name: "3"},
	}
	cels := []internal.Cel{
		{}, {}, {}, {},
	}
	result := output.SortLayers(layers, cels)
	if len(result) != 4 {
		t.Fatal("expected 4 layers")
	}
	if result[0].Name != "0" {
		t.Fatal("expected 0")
	}
	if result[1].Name != "1" {
		t.Fatal("expected 1")
	}
	if result[2].Name != "2" {
		t.Fatal("expected 2")
	}
	if result[3].Name != "3" {
		t.Fatal("expected 3")
	}
}

func TestCreateLayerImage(t *testing.T) {
	file := internal.File{
		Width: 1, Height: 1,
		ColorDepth: 32,
	}
	cel := internal.Cel{
		X: 0, Y: 0,
		Image: internal.CelImage{
			Width: 1, Height: 1,
			Bytes: []byte{0, 255, 0, 255},
		},
	}
	result := output.CreateCelImage(file, cel, 0)
	if result.Bounds().Dx()*result.Bounds().Dy() != 1 {
		t.Fatal("expected 1 pixel")
	}
	if matchesColor(result.At(0, 0), 0, 255, 0, 255) {
		t.Fatal("expected green pixel")
	}
}

func matchesColor(actual color.Color, r, g, b, a byte) bool {
	ar, ag, ab, aa := actual.RGBA()
	er, eg, eb, ea := color.NRGBA{r, g, b, a}.RGBA()
	return ar == er && ag == eg && ab == eb && aa == ea
}

func TestCreateFrameImage(t *testing.T) {
	file := internal.File{
		Width:      2,
		Height:     2,
		ColorDepth: 32,
		FileChunks: internal.FileChunks{
			Layers: []internal.Layer{
				{
					Flags:     internal.LayerFlagVisible,
					BlendMode: internal.BlendModeNormal,
					Opacity:   255,
				}, {
					Flags:     internal.LayerFlagVisible,
					BlendMode: internal.BlendModeNormal,
					Opacity:   255,
				}, {
					Flags:     internal.LayerFlagVisible,
					BlendMode: internal.BlendModeNormal,
					Opacity:   255,
				}, {
					Flags:     internal.LayerFlagVisible,
					BlendMode: internal.BlendModeNormal,
					Opacity:   255,
				},
			},
		},
	}
	frame := internal.Frame{
		Cels: []internal.Cel{
			{
				X: 0, Y: 0,
				Image: internal.CelImage{
					Width: 2, Height: 2,
					Bytes: []byte{
						255, 0, 0, 255,
						255, 0, 0, 255,
						255, 0, 0, 255,
						255, 0, 0, 255,
					},
				},
			},
			{
				X: 1, Y: 0,
				Image: internal.CelImage{
					Width: 1, Height: 1,
					Bytes: []byte{0, 255, 0, 255},
				},
			},
			{
				X: 0, Y: 1,
				Image: internal.CelImage{
					Width: 1, Height: 1,
					Bytes: []byte{0, 0, 255, 255},
				},
			},
			{
				X: 1, Y: 1,
				Image: internal.CelImage{
					Width: 1, Height: 1,
					Bytes: []byte{0, 0, 0, 255},
				},
			},
		},
	}
	image, err := output.CreateFrameImage(file, frame)
	if err != nil {
		t.Fatal(err)
	}
	if image.Bounds().Dx() != 2 && image.Bounds().Dy() != 2 {
		t.Fatal("expected 2x2 image")
	}
	if matchesColor(image.At(0, 0), 255, 0, 0, 255) {
		t.Fatal("expected red pixel")
	}
	if matchesColor(image.At(1, 0), 0, 255, 0, 255) {
		t.Fatal("expected green pixel")
	}
	if matchesColor(image.At(0, 1), 0, 0, 255, 255) {
		t.Fatal("expected blue pixel")
	}
	if matchesColor(image.At(1, 1), 0, 0, 0, 255) {
		t.Fatal("expected black pixel")
	}
}
