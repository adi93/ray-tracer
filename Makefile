PKGS := $(shell go list ./...)

build: test
	go build
	./ray-tracer
	open test.ppm

.PHONY: test
test:
	go test $(PKGS)

.PHONY: clean
clean:
	rm ray-tracer
