build-alloc:
	go build -gcflags '-m -l'


benchmark:
	go test -bench . -benchmem


benchmark-view:
	go test -bench . -benchmem -trace=pointer_trace.out
	go tool trace pointer_trace.out