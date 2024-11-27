package ase

type RGBAPixel [4]byte
type GrayscalePixel [2]byte
type IndexedPixel byte

type String struct {
	Length uint16 `label:"Length"`
	Value  string `label:"Value"`
}
