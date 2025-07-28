build: build-static
	goreleaser build --clean --skip validate -f .goreleaser.yaml

test:
	go test ./...

build-static:
	./scripts/build-static.sh
