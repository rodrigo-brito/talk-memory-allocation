package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	_ = NoAlloc()
	value := WithAlloc()
	_ = WithParameterPointer(value)
	_ = WithAllocStruct()
	_ = WithoutAllocStruct()
	WithInterface(&bytes.Buffer{})
	SliceDynamic()
	SliceWithCap()
}

func NoAlloc() float64 {
	x := 4.0
	y := x + 1

	return y
}

func WithAlloc() *float64 {
	x := 4.0
	y := x + 1

	return &y
}

type Float struct {
	float64
}

func WithoutAllocStruct() Float {
	x := 4.0
	y := x + 1
	return Float{y}
}

func WithAllocStruct() *Float {
	x := 4.0
	y := x + 1
	return &Float{y}
}

func WithParameterPointer(x *float64) float64 {
	y := *x + 1
	return y
}

func WithInterface(w io.Writer) {
	x := 1.0
	fmt.Fprint(w, x)
}

func SliceDynamic() {
	data := make([]int, 0)
	for i := 0; i < 1000; i++ {
		data = append(data, i)
	}
}

func SliceWithCap() {
	data := make([]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
}
