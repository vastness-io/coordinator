#!/bin/bash
for GOOS in darwin linux windows; do
	for GOARCH in 386 amd64; do
        echo "Building binary for $GOOS on $GOARCH"
        GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/${GOOS}/${GOARCH}/coordinator -ldflags "-X main.commit=${COMMIT} -X main.version=${VERSION}" cmd/coordinator/main.go
    done
done
echo "Complete"
