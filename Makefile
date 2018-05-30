all: build

build:
	go run tools\packr.go -v -z
	go build -ldflags "-X main.version=`cat VERSION`"

clean:
	go run tools\packr.go clean
	go clean

run:
	go run main.go -d 3s
