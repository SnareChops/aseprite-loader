package internal

import "image"

type Tileset struct {
	ID                uint32
	Name              string
	Flags             uint32
	Width             uint16
	Height            uint16
	Index             int16
	ExternalID        uint32
	ExternalTilesetID uint32
	Tiles             []Tile
	UserData          UserData
}

type Tile struct {
	image.Image
	UserData UserData
}
