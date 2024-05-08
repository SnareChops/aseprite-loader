package ase

type UserDataChunk struct {
	Flags uint32 `label:"Flags"`
}
type UserDataPropertyMaps struct {
	Size  uint32 `label:"Size in bytes of all properties maps stored in this chunk"`
	Count uint32 `label:"Number of properties maps"`
}
type UserDataPropertyMap struct {
	Key   uint32 `label:"Property map key"`
	Count uint32 `label:"Number of properties"`
}

type UserDataProperty struct {
	Name String `label:"Name"`
	Type uint16 `label:"Type"`
}
type Point struct {
	X int32 `label:"X coordinate value"`
	Y int32 `label:"Y coordinate value"`
}

type Size struct {
	Width  int32 `label:"Width value"`
	Height int32 `label:"Height value"`
}

type Rect struct {
	Origin Point `label:"Origin coordinates"`
	Size   Size  `label:"Rectangle size"`
}
type VectorHeader struct {
	Size uint32 `label:"Number of elements"`
	Type uint16 `label:"Element's type"`
}
type Element struct {
	Type  uint16 `label:"Element type"`
	Value []byte `label:"Element value"`
}

type NestedPropertyMap struct {
	Count uint32 `label:"Number of properties"`
	Data  []byte `label:"Nested properties data"`
}
