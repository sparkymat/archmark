all: archmark-linux archmark-worker-linux

archmark-linux: archmark.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o archmark-linux archmark.go

archmark-worker-linux: worker/cmd/archmark_worker.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o archmark-worker-linux worker/cmd/archmark_worker.go

docker: archmark-linux
	docker build . -t sparkymat/archmark

docker-worker: archmark-worker-linux
	docker build . -f Dockerfile_worker -t sparkymat/archmark-worker

clean:
	rm -rf archmark-linux archmark-worker-linux
