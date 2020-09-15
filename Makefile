BUILD_IMAGE := golang:1.15
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))
clean:
	rm -rf ./out/*
build:
	mkdir -p out/${OS}/${ARCH}
	docker run --rm -v $$(pwd):/src -w /src \
		 -e GOOS=${OS} -e GOARCH=${ARCH} -e GOARM=${GOARM}\
		 ${BUILD_IMAGE} go build -v -o out/${OS}/${ARCH} ./...
build-pi: OS := "linux"
build-pi: ARCH := "arm"
build-pi: GOARM := 7
build-pi: build
