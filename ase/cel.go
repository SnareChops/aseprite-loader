package ase

type CelChunk struct {
	Index     uint16  `label:"Layer Index"`
	XPosition int16   `label:"X Position"`
	YPosition int16   `label:"Y Position"`
	Opacity   byte    `label:"Opacity"`
	Type      uint16  `label:"Cel Type"`
	ZIndex    int16   `label:"Z-Index"`
	_         [5]byte `label:"For future use"`
}

type RawImageCel struct {
	Width  uint16 `label:"Width"`
	Height uint16 `label:"Height"`
}

type LinkedCel struct {
	LinkFrame uint16 `label:"Frame position to link with"`
}

type CompressedImageCel struct {
	Width  uint16 `label:"Width"`
	Height uint16 `label:"Height"`
}

type CompressedTilemapCel struct {
	Width       uint16   `label:"Width in number of tiles"`
	Height      uint16   `label:"Height in number of tiles"`
	BitsPerTile uint16   `label:"Bits per tile (at the moment it's always 32-bit per tile)"`
	TileIDMask  uint32   `label:"Bitmask for tile ID"`
	XFlipMask   uint32   `label:"Bitmask for X flip"`
	YFlipMask   uint32   `label:"Bitmask for Y flip"`
	DFlipMask   uint32   `label:"Bitmask for diagonal flip"`
	_           [10]byte `label:"Reserved"`
}

type RawImage[T IndexedPixel | GrayscalePixel | RGBAPixel] struct {
	RawImageCel
	Data []T `label:"Pixel data"`
}

type CompressedImage struct {
	CompressedImageCel
	Data []byte `label:"Compressed image data"`
}

type CompressedTilemap struct {
	CompressedTilemapCel
	Data []byte `label:"Compressed tilemap data"`
}
