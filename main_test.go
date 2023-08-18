package main

import (
	"bytes"
	"math/rand"
	"testing"
)

func BenchmarkNoAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NoAlloc()
	}
	b.ReportAllocs()
}

func BenchmarkWithoutAllocStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = WithoutAllocStruct()
	}
	b.ReportAllocs()
}

func BenchmarkWithAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = WithAlloc()
	}
	b.ReportAllocs()
}

func BenchmarkWithParameterPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := rand.Float64()
		_ = WithParameterPointer(&x)
	}
	b.ReportAllocs()
}

func BenchmarkWithInterface(b *testing.B) {
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		WithInterface(w)
	}
	b.ReportAllocs()
}

func BenchmarkWithAllocStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = WithAllocStruct()
	}
	b.ReportAllocs()
}

func BenchmarkSliceDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceDynamic()
	}
	b.ReportAllocs()
}

func BenchmarkSliceCapacity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceWithCap()
	}
	b.ReportAllocs()
}
