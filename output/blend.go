package output

import (
	"image"
	"image/color"

	"hawx.me/code/img/blend"
	"hawx.me/code/img/utils"
)

func Divide(a, b image.Image) image.Image {
	return blend.BlendPixels(a, b, func(c, d color.Color) color.Color {
		i, j, k, l := utils.RatioRGBA(c)
		m, n, o, p := utils.RatioRGBA(d)
		r := i / m
		g := j / n
		b := k / o
		a := l / p
		return color.NRGBA{
			uint8(utils.Truncatef(r * 255)),
			uint8(utils.Truncatef(g * 255)),
			uint8(utils.Truncatef(b * 255)),
			uint8(utils.Truncatef(a * 255)),
		}
	})
}
