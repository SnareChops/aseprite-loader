package internal

import "image"

type LayerFlag uint16

const (
	LayerFlagVisible LayerFlag = 1 << iota
	LayerFlagEditable
	LayerFlagLockMovement
	LayerFlagBackground
	LayerFlagPreferLinkedCels
	LayerFlagCollapsed
	LayerFlagReference
)

type BlendMode uint16

const (
	BlendModeNormal     BlendMode = 0
	BlendModeMultiply   BlendMode = 1
	BlendModeScreen     BlendMode = 2
	BlendModeOverlay    BlendMode = 3
	BlendModeDarken     BlendMode = 4
	BlendModeLighten    BlendMode = 5
	BlendModeColorDodge BlendMode = 6
	BlendModeColorBurn  BlendMode = 7
	BlendModeHardLight  BlendMode = 8
	BlendModeSoftLight  BlendMode = 9
	BlendModeDifference BlendMode = 10
	BlendModeExclusion  BlendMode = 11
	BlendModeHue        BlendMode = 12
	BlendModeSaturation BlendMode = 13
	BlendModeColor      BlendMode = 14
	BlendModeLuminosity BlendMode = 15
	BlendModeAddition   BlendMode = 16
	BlendModeSubtract   BlendMode = 17
	BlendModeDivide     BlendMode = 18
)

type Layer struct {
	Name          string
	Flags         LayerFlag
	Type          uint16
	DefaultWidth  uint16
	DefaultHeight uint16
	BlendMode     BlendMode
	Opacity       byte
	TilesetID     uint32
	UserData      UserData
	Image         image.Image
}
