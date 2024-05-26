package lib_test

import (
	"image"
	"testing"

	"github.com/SnareChops/aseprite-loader/lib"
)

func TestSlice(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 1280, 1280))
	result, err := lib.Slice(img, 128, 128)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 100 {
		t.Fatal("Expected 100 slices, got", len(result))
	}
}

func TestSmashAndSlice(t *testing.T) {
	layer := lib.Layer{
		Name:      "test",
		BlendMode: lib.BlendModeNormal,
		Opacity:   255,
		IsVisible: true,
		Image:     image.NewNRGBA(image.Rect(0, 0, 1280, 1280)),
	}
	result, err := lib.SmashAndSlice([]lib.Layer{layer}, 128, 128)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 100 {
		t.Fatal("Expected 100 slices, got", len(result))
	}
}
