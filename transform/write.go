package transform

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
)

func writeType(out io.Writer, label string, data []byte, typ reflect.Type) (err error) {
	log.Println("Writing type", typ, label)
	switch typ.Kind() {
	case reflect.Uint8:
		return writeLine(out, data, label, fmt.Sprint(uint8(data[0])), "uint8")
	case reflect.Int8:
		return writeLine(out, data, label, fmt.Sprint(int8(data[0])), "int8")
	case reflect.Uint16:
		return writeLine(out, data, label, fmt.Sprint(binary.LittleEndian.Uint16(data)), "uint16")
	case reflect.Int16:
		return writeLine(out, data, label, fmt.Sprint(int16(binary.LittleEndian.Uint16(data))), "int16")
	case reflect.Uint32:
		return writeLine(out, data, label, fmt.Sprint(binary.LittleEndian.Uint32(data)), "uint32")
	case reflect.Int32:
		return writeLine(out, data, label, fmt.Sprint(int32(binary.LittleEndian.Uint32(data))), "int32")
	case reflect.Uint64:
		return writeLine(out, data, label, fmt.Sprint(binary.LittleEndian.Uint64(data)), "uint64")
	case reflect.Int64:
		return writeLine(out, data, label, fmt.Sprint(int64(binary.LittleEndian.Uint64(data))), "int64")
	case reflect.Float32:
		return writeLine(out, data, label, fmt.Sprint(float32(binary.LittleEndian.Uint32(data))), "float32")
	case reflect.Float64:
		return writeLine(out, data, label, fmt.Sprint(float64(binary.LittleEndian.Uint64(data))), "float64")
	case reflect.String:
		return writeLine(out, data, label, string(data), "string")
	default:
		return writeBytes(out, label, data)
	}
}

func writeLine(out io.Writer, data []byte, label string, value string, typ string) (err error) {
	line := []string{}
	for _, b := range data {
		line = append(line, fmt.Sprintf("0x%02x", b))
	}
	comment := "//"
	if label != "" {
		comment += " " + label
	}
	if value != "" {
		comment += " " + value
	}
	if typ != "" {
		comment += " (" + typ + ")"
	}
	line = append(line, comment)
	_, err = out.Write(append([]byte(strings.Join(line, ", ")), '\n'))
	return
}

func writeBytes(out io.Writer, label string, data []byte) (err error) {
	line := []string{}
	for _, b := range data {
		line = append(line, fmt.Sprintf("0x%02x", b))
	}
	if label != "" {
		line = append(line, "// "+label)
	}
	_, err = out.Write(append([]byte(strings.Join(line, ", ")), '\n'))
	return
}
