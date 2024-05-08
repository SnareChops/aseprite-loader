package ase

type FileHeader struct {
	FileSize       uint32    `label:"File Size"`
	Magic          uint16    `label:"Magic Number"`
	Frames         uint16    `label:"Number of Frames"`
	Width          uint16    `label:"Width in pixels"`
	Height         uint16    `label:"Height in pixels"`
	ColorDepth     uint16    `label:"Color depth (bits per pixel)"`
	Flags          uint32    `label:"Flags"`
	Speed          uint16    `label:"Speed (deprecated)"`
	_              [2]uint32 `label:"Unused"`
	PaletteIndex   byte      `label:"Palette Entry Index"`
	_              [3]byte   `label:"Ignored"`
	NumberOfColors uint16    `label:"Number of colors"`
	PixelWidth     byte      `label:"Pixel Width"`
	PixelHeight    byte      `label:"Pixel Height"`
	XGridPosition  int16     `label:"X Grid Position"`
	YGridPosition  int16     `label:"Y Grid Position"`
	GridWidth      uint16    `label:"Grid Width"`
	GridHeight     uint16    `label:"Grid Height"`
	_              [84]byte  `label:"For future use"`
}
