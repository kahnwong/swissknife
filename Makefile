build: build-dynamic
	goreleaser build --clean --skip validate -f .goreleaser.yaml

test:
	go test ./...

build-dynamic:
	./scripts/build-dynamic.sh
