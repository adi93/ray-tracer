PKGS := $(shell go list ./...)

RAY_TRACER := ray-tracer
build: test
	go build -o $(RAY_TRACER)
	./$(RAY_TRACER)
	open test.ppm

.PHONY: test
test:
	go test $(PKGS)

.PHONY: clean
clean:
	go clean
	rm *.ppm
