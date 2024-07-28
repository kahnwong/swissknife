build:
	goreleaser release --skip publish --skip validate --skip archive --clean -f .goreleaser-linux.yaml

test:
	go test ./...
