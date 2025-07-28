build: build-dynamic
	goreleaser release --skip publish --skip validate --skip archive --clean -f .goreleaser-linux.yaml

test:
	go test ./...

build-dynamic:
	./scripts/build-dynamic.sh
