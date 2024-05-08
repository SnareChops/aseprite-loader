package internal

import "image/color"

type Tags []Tag

type Tag struct {
	Name          string
	From          uint16
	To            uint16
	LoopDirection byte
	Repeat        byte
	Color         color.Color
	UserData      UserData
}
