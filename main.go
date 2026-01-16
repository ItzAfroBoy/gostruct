// A go implementation of Python's Struct module
package gostruct

import (
	"encoding/binary"
	"math"
	"reflect"
	"strings"
)

// Unpack buffer using format string.
// Returns array containing parsed data
func Unpack(format string, buffer []byte) []any {
	var bin_func binary.ByteOrder
	var arr_ret []any
	var offset int

	arr := strings.Split(format, "")
	endian := arr[0]

	switch endian {
	case "<":
		bin_func = binary.LittleEndian
		arr = arr[1:]
	case ">":
		bin_func = binary.BigEndian
		arr = arr[1:]
	default:
		bin_func = binary.NativeEndian
	}

	for _, v := range arr {
		switch v {
		case "i":
			arr_ret = append(arr_ret, int32(bin_func.Uint32(buffer[offset:offset+4])))
			offset += 4
		case "I":
			arr_ret = append(arr_ret, bin_func.Uint32(buffer[offset:offset+4]))
			offset += 4
		case "f":
			arr_ret = append(arr_ret, math.Float32frombits(bin_func.Uint32(buffer[offset:offset+4])))
			offset += 4
		case "h":
			arr_ret = append(arr_ret, int16(bin_func.Uint16(buffer[offset:offset+2])))
			offset += 2
		case "H":
			arr_ret = append(arr_ret, bin_func.Uint16(buffer[offset:offset+2]))
			offset += 2
		case "b":
			arr_ret = append(arr_ret, int8(buffer[offset]))
			offset += 1
		case "B":
			arr_ret = append(arr_ret, uint8(buffer[offset]))
			offset += 1
		}
	}

	return arr_ret
}

// Unpack buffer using format string.
// Populates v using reflect. 
// Buffer contents must match amount of struct fields
func UnpackToStruct(format string, buffer []byte, v any) {
	unpacked := Unpack(format, buffer)
	e := reflect.ValueOf(v).Elem()
	vf := reflect.VisibleFields(e.Type())
	for i := range vf {
		if len(unpacked) != len(vf) {
			panic("Mismatched size")
		}

		if reflect.ValueOf(unpacked[i]).Type() == e.Field(i).Type() {
			e.Field(i).Set(reflect.ValueOf(unpacked[i]))
		}
	}
}
