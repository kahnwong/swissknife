build:
	./scripts/build-static.sh
build-test:
	goreleaser -f .goreleaser-linux-amd64.yaml build --skip validate --snapshot --clean
test:
	go test ./...
