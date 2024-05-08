package internal

type CelType uint16

const (
	CelTypeRawImage CelType = iota
	CelTypeLinkedCel
	CelTypeCompressedImage
	CelTypeCompressedTilemap
)

type Cel struct {
	LayerIndex uint16
	Type       CelType
	Opacity    byte
	X          int
	Y          int
	Width      uint16
	Height     uint16
	ZIndex     int16
	Image      CelImage
	Link       uint16
	Tilemap    CelTilemap
	Raw        []byte
	CelExtra
	UserData UserData
}

type CelImage struct {
	Width  int
	Height int
	Bytes  []byte
}

type CelTilemap struct {
	Width       uint16
	Height      uint16
	BitsPerTile uint16
	TileIDMask  uint32
	XFlipMask   uint32
	YFlipMask   uint32
	DFlipMask   uint32
	Tiles       []uint32
	// Bytes       []byte
}
