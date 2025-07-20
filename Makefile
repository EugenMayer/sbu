init:
	go mod tidy
	go mod verify

update:
	go get -u
	go mod tidy

tests:
	go test ./...

build:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags netgo -o dist/sbu-macos-amd64 main.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -tags netgo -o dist/sbu-macos-arm64 main.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o dist/sbu-linux-amd64 main.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags netgo -o dist/sbu-windows-amd64 main.go

build-dev:
	go build -o dist/sbu main.go
