build:
	./scripts/build-static.sh

test:
	go test ./...

build-darwin:
	CGO_ENABLED=1 go build -ldflags="-s -w -linkmode 'external'" .
