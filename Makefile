all: archmark


archmark-linux: archmark.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o archmark-linux archmark.go

docker: archmark-linux
	docker build . -t sparkymat/archmark

clean:
	rm -rf archmark-linux
