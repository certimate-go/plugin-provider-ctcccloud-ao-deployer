.PHONY: build test

build:
	GOWORK=off go build -mod=readonly -o ctcccloud-ao.exe .

test:
	GOWORK=off go test -mod=readonly ./...
