# talk-memory-allocation

Discussion about memory allocation in Go (Stack / Heap)

## Commands

- `make build-alloc` - build with alloc information
- `make benchmark ` - run benchmark
- `make benchmark-view` - run benchmark and view trace result


## Build Alloc

```
go build -gcflags '-m -l'
# github.com/rodrigo-brito/talk-memory-allocation
./main.go:29:2: moved to heap: y
./main.go:50:27: x does not escape
./main.go:47:9: &Float{...} escapes to heap
./main.go:55:20: leaking param: w
./main.go:57:12: ... argument does not escape
./main.go:57:16: x escapes to heap
./main.go:61:14: make([]int, 0) does not escape
./main.go:68:14: make([]int, 1000, 1000) does not escape
./main.go:15:16: &bytes.Buffer{} escapes to heap
```

## Results
```
BenchmarkNoAlloc-8                      1000000000               0.3230 ns/op          0 B/op          0 allocs/op
BenchmarkWithoutAllocStruct-8           1000000000               0.3137 ns/op          0 B/op          0 allocs/op
BenchmarkWithAlloc-8                    1000000000               0.3137 ns/op          0 B/op          0 allocs/op
BenchmarkWithParameterPointer-8         317108812                3.778 ns/op           0 B/op          0 allocs/op
BenchmarkWithInterface-8                15643008                76.01 ns/op           10 B/op          1 allocs/op
BenchmarkWithAllocStruct-8              1000000000               0.3131 ns/op          0 B/op          0 allocs/op
BenchmarkSliceDynamic-8                   497601              2321 ns/op           25208 B/op         12 allocs/op
BenchmarkSliceWithSize-8                 2433013               493.2 ns/op             0 B/op          0 allocs/op
BenchmarkSliceCapacity-8                 3717766               327.1 ns/op             0 B/op          0 allocs/op
```