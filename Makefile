NAME := awssecret2env
MAIN_SRC := cmd/$(NAME)/main.go
LATEST_TAG := $(shell git describe --tag)

.PHONY: default build clean install build-all checkout-latest-tag upload-master upload-release env release master

default: build

build:
	go build -o build/bin/$(NAME) $(MAIN_SRC)

clean:
	go clean
	rm -rf build/*
	mkdir -p build/bin
	touch build/bin/.gitkeep

install: build
	cp ./build/bin/$(NAME) /usr/local/bin/$(NAME)

build-all: clean build
	mkdir -p build/dist
	GOOS=darwin GOARCH=amd64 go build -o build/dist/$(NAME)-macos $(MAIN_SRC)
	GOOS=darwin GOARCH=arm64 go build -o build/dist/$(NAME)-macos-arm64 $(MAIN_SRC)
	GOOS=windows GOARCH=amd64 go build -o build/dist/$(NAME)-windows $(MAIN_SRC)
	GOOS=linux GOARCH=amd64 go build -o build/dist/$(NAME)-linux64 $(MAIN_SRC)
	GOOS=linux GOARCH=arm GOARM=7 go build -o build/dist/$(NAME)-linuxarm7 $(MAIN_SRC)
	GOOS=linux GOARCH=arm GOARM=6 go build -o build/dist/$(NAME)-linuxarm6 $(MAIN_SRC)

checkout-latest-tag:
	git checkout $(LATEST_TAG)

upload-master:
	aws s3 cp build/dist/$(NAME)-macos s3://awssecret2env/master/$(NAME)-macos --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-macos-arm64 s3://awssecret2env/master/$(NAME)-macos-arm64 --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-windows s3://awssecret2env/master/$(NAME)-windows --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-linux64 s3://awssecret2env/master/$(NAME)-linux64 --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-linuxarm7 s3://awssecret2env/master/$(NAME)-linuxarm7 --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-linuxarm6 s3://awssecret2env/master/$(NAME)-linuxarm6 --content-type application/octet-stream

upload-release:
	aws s3 cp build/dist/$(NAME)-macos s3://awssecret2env/$(LATEST_TAG)/$(NAME)-macos --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-macos-arm64 s3://awssecret2env/$(LATEST_TAG)/$(NAME)-macos-arm64 --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-windows s3://awssecret2env/$(LATEST_TAG)/$(NAME)-windows --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-linux64 s3://awssecret2env/$(LATEST_TAG)/$(NAME)-linux64 --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-linuxarm7 s3://awssecret2env/$(LATEST_TAG)/$(NAME)-linuxarm7 --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-linuxarm6 s3://awssecret2env/$(LATEST_TAG)/$(NAME)-linuxarm6 --content-type application/octet-stream

	aws s3 cp s3://awssecret2env/$(LATEST_TAG)/$(NAME)-macos s3://awssecret2env/latest/$(NAME)-macos
	aws s3 cp s3://awssecret2env/$(LATEST_TAG)/$(NAME)-macos-arm64 s3://awssecret2env/latest/$(NAME)-macos-arm64
	aws s3 cp s3://awssecret2env/$(LATEST_TAG)/$(NAME)-windows s3://awssecret2env/latest/$(NAME)-windows
	aws s3 cp s3://awssecret2env/$(LATEST_TAG)/$(NAME)-linux64 s3://awssecret2env/latest/$(NAME)-linux64
	aws s3 cp s3://awssecret2env/$(LATEST_TAG)/$(NAME)-linuxarm7 s3://awssecret2env/latest/$(NAME)-linuxarm7
	aws s3 cp s3://awssecret2env/$(LATEST_TAG)/$(NAME)-linuxarm6 s3://awssecret2env/latest/$(NAME)-linuxarm6

env: build
	./build/bin/$(NAME) --export --output .env secrets.txt

release: checkout-latest-tag build-all upload-release

master: build-all upload-master
