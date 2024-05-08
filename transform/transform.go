package transform

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/SnareChops/aseprite-loader/trace"
)

func transform[T any](in io.Reader, out io.Writer) (data T, err error) {
	trace.Log("transform", reflect.TypeOf(data).Name())
	err = binary.Read(in, binary.LittleEndian, &data)
	if err != nil {
		return
	}
	err = write(out, data)
	return
}

func write[T any](out io.Writer, data T) (err error) {
	trace.Log("write", reflect.TypeOf(data).Name(), out)
	if out == nil || (reflect.ValueOf(out).Kind() == reflect.Ptr && reflect.ValueOf(out).IsNil()) {
		return
	}
	typ := reflect.TypeOf(data)
	switch typ.Kind() {
	case reflect.Struct:
		out.Write([]byte("// " + typ.Name() + "\n"))
		for i := range typ.NumField() {
			buffer := new(bytes.Buffer)
			field := typ.Field(i)
			label := field.Tag.Get("label")
			var val []byte
			val, err = getBytes(reflect.ValueOf(data).Field(i))
			if err != nil {
				panic(err)
			}
			err = binary.Write(buffer, binary.LittleEndian, val)
			if err != nil {
				panic(err)
			}
			err = writeBytes(out, label, buffer.Bytes())
		}
	default:
		buffer := new(bytes.Buffer)
		err = binary.Write(buffer, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
		err = writeBytes(out, "", buffer.Bytes())
	}
	return
}

func writeBytes(out io.Writer, label string, data []byte) (err error) {
	line := []string{}
	for _, b := range data {
		line = append(line, fmt.Sprintf("0x%02x", b))
	}
	if label != "" {
		line = append(line, "// "+label)
	} else {
		line = append(line, "")
	}
	_, err = out.Write(append([]byte(strings.Join(line, ", ")), '\n'))
	return
}

func getBytes(v reflect.Value) (value []byte, err error) {
	trace.Log("getBytes", v.Kind().String())
	buffer := new(bytes.Buffer)
	switch v.Kind() {
	case reflect.Uint8:
		err = binary.Write(buffer, binary.LittleEndian, uint8(v.Uint()))
	case reflect.Uint16:
		err = binary.Write(buffer, binary.LittleEndian, uint16(v.Uint()))
	case reflect.Uint32:
		err = binary.Write(buffer, binary.LittleEndian, uint32(v.Uint()))
	case reflect.Uint64:
		err = binary.Write(buffer, binary.LittleEndian, v.Uint())
	case reflect.Int8:
		err = binary.Write(buffer, binary.LittleEndian, int8(v.Int()))
	case reflect.Int16:
		err = binary.Write(buffer, binary.LittleEndian, int16(v.Int()))
	case reflect.Int32:
		err = binary.Write(buffer, binary.LittleEndian, int32(v.Int()))
	case reflect.Int64:
		err = binary.Write(buffer, binary.LittleEndian, v.Int())
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			var val []byte
			val, err = getBytes(v.Index(i))
			if err != nil {
				return
			}
			_, err = buffer.Write(val)
		}
	case reflect.String:
		err = binary.Write(buffer, binary.LittleEndian, []byte(v.String()))
	default:
		err = fmt.Errorf("unsupported type %s", v.Kind().String())
	}
	if err != nil {
		return
	}
	value = buffer.Bytes()
	return
}
