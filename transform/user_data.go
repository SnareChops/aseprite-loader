package transform

import (
	"errors"
	"image/color"
	"io"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func transformUserData(in io.Reader, out io.Writer) (data internal.UserData, err error) {
	trace.Log("transformUserData")
	userData, err := transform[ase.UserDataChunk](in, out)
	if err != nil {
		return
	}
	data.Flags = userData.Flags
	if userData.Flags&0x1 == 0x1 {
		var text ase.String
		text, err = transformString(in, out)
		if err != nil {
			return
		}
		data.Text = text.Value
	}
	if userData.Flags&0x2 == 0x2 {
		var clr [4]byte
		clr, err = transform[[4]byte](in, out)
		if err != nil {
			return
		}
		data.Color = color.NRGBA{clr[0], clr[1], clr[2], clr[3]}
	}
	if userData.Flags&0x4 == 0x4 {
		data.Properties = internal.PropertyMap{}
		var header ase.UserDataPropertyMaps
		header, err = transform[ase.UserDataPropertyMaps](in, out)
		if err != nil {
			return
		}
		for range header.Count {
			var key uint32
			var properties []internal.Property
			key, properties, err = transformPropertyMap(in, out)
			if err != nil {
				return
			}
			data.Properties[key] = properties
		}
	}
	return
}

func transformPropertyMap(in io.Reader, out io.Writer) (key uint32, data []internal.Property, err error) {
	trace.Log("transformPropertyMap")
	var propertyMap ase.UserDataPropertyMap
	propertyMap, err = transform[ase.UserDataPropertyMap](in, out)
	if err != nil {
		return
	}
	key = propertyMap.Key
	for range propertyMap.Count {
		var property internal.Property
		property, err = transformProperty(in, out)
		if err != nil {
			return
		}
		data = append(data, property)
	}
	return
}

func transformProperty(in io.Reader, out io.Writer) (data internal.Property, err error) {
	trace.Log("transformProperty")
	var name ase.String
	name, err = transformString(in, out)
	if err != nil {
		return
	}
	data.Name = name.Value
	var typ uint16
	typ, err = transform[uint16](in, out)
	if err != nil {
		return
	}
	data.Type = typ
	data.Element, err = transformElement(in, out, data.Type)
	return
}

func transformElement(in io.Reader, out io.Writer, typ uint16) (data internal.Element, err error) {
	trace.Log("transformElement")
	if typ == 0 {
		typ, err = transform[uint16](in, out)
		if err != nil {
			return
		}
	}
	switch typ {
	case 0x01:
		var b byte
		b, err = transform[byte](in, out)
		data.Bool = true
		if b == 0 {
			data.Bool = false
		}
	case 0x02:
		var i int8
		i, err = transform[int8](in, out)
		data.Int = int(i)
	case 0x03:
		var u uint8
		u, err = transform[uint8](in, out)
		data.Uint = uint(u)
	case 0x04:
		var i int16
		i, err = transform[int16](in, out)
		data.Int = int(i)
	case 0x05:
		var u uint16
		u, err = transform[uint16](in, out)
		data.Uint = uint(u)
	case 0x06:
		var i int32
		i, err = transform[int32](in, out)
		data.Int = int(i)
	case 0x07:
		var u uint32
		u, err = transform[uint32](in, out)
		data.Uint = uint(u)
	case 0x08:
		var i int64
		i, err = transform[int64](in, out)
		data.Int = int(i)
	case 0x09:
		var u uint64
		u, err = transform[uint64](in, out)
		data.Uint = uint(u)
	case 0x0a:
		data.Fixed, err = transform[[4]byte](in, out)
	case 0x0b:
		var f float32
		f, err = transform[float32](in, out)
		data.Float = float64(f)
	case 0x0c:
		data.Float, err = transform[float64](in, out)
	case 0x0d:
		var val ase.String
		val, err = transformString(in, out)
		data.String = val.Value
	case 0x0e:
		var val ase.Point
		val, err = transform[ase.Point](in, out)
		data.Point = internal.Point{X: val.X, Y: val.Y}
	case 0x0f:
		var val ase.Size
		val, err = transform[ase.Size](in, out)
		data.Size = internal.Size{Width: val.Width, Height: val.Height}
	case 0x10:
		var val ase.Rect
		val, err = transform[ase.Rect](in, out)
		data.Rect = internal.Rect{
			Origin: internal.Point{X: val.Origin.X, Y: val.Origin.Y},
			Size:   internal.Size{Width: val.Size.Width, Height: val.Size.Height},
		}
	case 0x11:
		data.Vector, err = transformVector(in, out)
	case 0x12:
		data.Properties, err = transformNestedPropertyMap(in, out)
	case 0x013:
		data.UUID, err = transform[[16]byte](in, out)
	default:
		err = errors.New("transformElement: unknown element type")
	}
	return
}

func transformVector(in io.Reader, out io.Writer) (data []internal.Element, err error) {
	trace.Log("transformVector")
	var header ase.VectorHeader
	header, err = transform[ase.VectorHeader](in, out)
	if err != nil {
		return
	}
	for range header.Size {
		var element internal.Element
		element, err = transformElement(in, out, header.Type)
		if err != nil {
			return
		}
		data = append(data, element)
	}
	return
}

func transformNestedPropertyMap(in io.Reader, out io.Writer) (data []internal.Property, err error) {
	trace.Log("transformNestedPropertyMap")
	var count uint32
	count, err = transform[uint32](in, out)
	if err != nil {
		return
	}
	for range count {
		var property internal.Property
		property, err = transformProperty(in, out)
		if err != nil {
			return
		}
		data = append(data, property)
	}
	return
}

// func transformElements(in io.Reader, out io.Writer, size uint32, typ uint16) (data []internal.Element, err error) {
// 	for range size {
// 		var element internal.Element
// 		element, err = transformElement(in, out, typ)
// 		if err != nil {
// 			return
// 		}
// 		data = append(data, element)
// 	}
// 	return
// }
