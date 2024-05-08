package internal

import "image/color"

type Palette struct {
	FirstIndex uint32
	LastIndex  uint32
	Entries    []PaletteEntry
	UserData   UserData
}

type PaletteEntry struct {
	Name  string
	Flags uint16
	Color color.Color
}
