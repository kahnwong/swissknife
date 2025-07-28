build: build-static
	goreleaser build --clean --skip validate -f .goreleaser-linux-amd64.yaml

test:
	go test ./...

build-static:
	./scripts/build-static.sh
