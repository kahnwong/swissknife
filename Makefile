build-lib:
	./scripts/build-lib.sh
build:
	RELEASE_GOOS=linux RELEASE_GOARCH=amd64 TARGET=x86_64-unknown-linux-gnu CC=gcc goreleaser build --snapshot --clean
test:
	go test ./...
