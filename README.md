# talk-memory-allocation

Discussion about memory allocation in Go (Stack / Heap)

## Commands

- `make build-alloc` - build with alloc information
- `make benchmark ` - run benchmark
- `make benchmark-view` - run benchmark and view trace result

## Results
```
BenchmarkNoAlloc-8                      1000000000               0.3188 ns/op          0 B/op          0 allocs/op
BenchmarkWithoutAllocStruct-8           1000000000               0.3134 ns/op          0 B/op          0 allocs/op
BenchmarkWithAlloc-8                    1000000000               0.3130 ns/op          0 B/op          0 allocs/op
BenchmarkWithParameterPointer-8         319199694                3.760 ns/op           0 B/op          0 allocs/op
BenchmarkWithInterface-8                15736629                75.58 ns/op           10 B/op          1 allocs/op
BenchmarkWithAllocStruct-8              1000000000               0.3130 ns/op          0 B/op          0 allocs/op
BenchmarkSliceDynamic-8                   489392              2294 ns/op           25208 B/op         12 allocs/op
BenchmarkSliceCapacity-8                 3724827               322.2 ns/op             0 B/op          0 allocs/op
```