package internal

import "image/color"

type Tags []Tag

type Tag struct {
	Name          string
	From          uint16
	To            uint16
	LoopDirection byte
	Repeat        uint16
	Color         color.Color
	UserData      UserData
}
