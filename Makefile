build: build-static
	goreleaser release --skip publish --skip validate --skip archive --clean -f .goreleaser-linux.yaml

test:
	go test ./...

build-static:
	./scripts/build-static.sh
