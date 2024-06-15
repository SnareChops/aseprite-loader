package internal

import "image/color"

type File struct {
	Width          int
	Height         int
	ColorDepth     uint16
	PaletteIndex   byte
	NumberOfColors uint16
	PixelWidth     byte
	PixelHeight    byte
	XGridPosition  int16
	YGridPosition  int16
	GridWidth      uint16
	GridHeight     uint16
	Frames         []Frame
	FileChunks
}

type FileChunks struct {
	Palette       []color.Color
	ExternalFiles ExternalFiles
	ColorProfile  ColorProfile
	Tags          Tags
	Slices        []Slice
	Tilesets      []Tileset
	Layers        []Layer
}
