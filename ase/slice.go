package ase

type SliceChunk struct {
	Keys  uint32 `label:"Number of slice keys"`
	Flags uint32 `label:"Flags"`
	_     uint32 `label:"Reserved"`
}

type SliceKey struct {
	Frame  uint32 `label:"Frame number"`
	X      int32  `label:"Slice X origin coordinate in the sprite"`
	Y      int32  `label:"Slice Y origin coordinate in the sprite"`
	Width  uint32 `label:"Slice width"`
	Height uint32 `label:"Slice height"`
}

type Slice9Patches struct {
	X      int32  `label:"Center X position (relative to slice bounds)"`
	Y      int32  `label:"Center Y position"`
	Width  uint32 `label:"Center width"`
	Height uint32 `label:"Center height"`
}

type SlicePivot struct {
	X int32 `label:"Pivot X position (relative to the slice origin)"`
	Y int32 `label:"Pivot Y position"`
}
