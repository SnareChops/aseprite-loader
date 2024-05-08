package internal

type Slice struct {
	Name     string
	Flags    uint32
	Keys     []SliceKey
	UserData UserData
}

type SliceKey struct {
	Frame   uint32
	X, Y    int32
	Width   uint32
	Height  uint32
	Patches Slice9Patches
	Pivot   SlicePivot
}

type Slice9Patches struct {
	X, Y   int32
	Width  uint32
	Height uint32
}

type SlicePivot struct {
	X, Y int32
}
