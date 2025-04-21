.PHONY: run
run: # Run `task`
	go run ./cmd/task

.PHONY: build
build: # Build `task` native binary
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -tags timetzdata -o ./build/task ./cmd/task
