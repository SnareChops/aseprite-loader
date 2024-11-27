package internal

import "image/color"

type UserData struct {
	Flags      uint32
	Text       string
	Color      color.Color
	Properties PropertyMap
}

type PropertyMap map[uint32][]Property

type Property struct {
	Name string
	Type uint16
	Element
}

type Element struct {
	Bool       bool
	Int        int
	Uint       uint
	Float      float64
	Fixed      [4]byte
	String     string
	Point      Point
	Size       Size
	Rect       Rect
	Vector     []Element
	Properties []Property
	UUID       [16]byte
}

type Point struct {
	X, Y int32
}

type Size struct {
	Width, Height int32
}

type Rect struct {
	Origin Point
	Size   Size
}
