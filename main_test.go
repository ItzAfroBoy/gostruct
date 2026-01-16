package gostruct_test

import (
	"fmt"
	"github.com/ItzAfroBoy/gostruct"
)

type Data struct {
	A int32
	B int32
	C float32
	D float32
	E uint8
	F uint32
}

func Example() {
	byteStream := receive() // Data to unpack
	packet := Data{}
	gostruct.UnpackToStruct("<iiffBI", byteStream, &packet)
	fmt.Println(packet)
}
