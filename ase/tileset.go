package ase

type TilesetChunk struct {
	ID     uint32   `label:"Tileset ID"`
	Flags  uint32   `label:"Tileset flags"`
	Tiles  uint32   `label:"Number of tiles"`
	Width  uint16   `label:"Tile width"`
	Height uint16   `label:"Tile height"`
	Index  int16    `label:"Base index"`
	_      [14]byte `label:"Reserved"`
}

type TilesetExternalData struct {
	ID        uint32 `label:"ID of the external file"`
	TilesetID uint32 `label:"Tileset ID in the external file"`
}

type TilesetCompressedData struct {
	Length uint32 `label:"Compressed data length"`
	Image  []byte `label:"Compressed image data"`
}
