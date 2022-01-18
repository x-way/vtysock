linux-binaries:
	GOOS=linux GOARCH=amd64 go build -o vtysock.amd64 .
	GOOS=linux GOARCH=arm64 go build -o vtysock.aarch64 .
