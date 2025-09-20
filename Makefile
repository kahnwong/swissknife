build:
	./scripts/build-static.sh

test:
	go test ./...

build-darwin:
	./scripts/build-static-ci-darwin-arm64.sh
	CGO_ENABLED=1 go build -ldflags="-s -w -linkmode 'external'" .
