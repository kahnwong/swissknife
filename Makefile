build:
	goreleaser release --skip publish --skip validate --clean -f .goreleaser-linux.yaml

# --------
test:
	go test ./...
