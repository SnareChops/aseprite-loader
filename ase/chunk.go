package ase

const (
	ChunkTypeOldPalette    = 0x0004
	ChunkTypeAltPalette    = 0x0011
	ChunkTypeLayer         = 0x2004
	ChunkTypeCel           = 0x2005
	ChunkTypeCelExtra      = 0x2006
	ChunkTypeColorProfile  = 0x2007
	ChunkTypeExternalFiles = 0x2008
	ChunkTypeMask          = 0x2016
	ChunkTypeTags          = 0x2018
	ChunkTypePalette       = 0x2019
	ChunkTypeUserData      = 0x2020
	ChunkTypeSlice         = 0x2022
	ChunkTypeTileset       = 0x2023
)

type ChunkHeader struct {
	ChunkSize uint32 `label:"Chunk Size"`
	ChunkType uint16 `label:"Chunk Type"`
}
