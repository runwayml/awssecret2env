NAME := awssecret2env
MAIN_SRC := cmd/$(NAME)/main.go

.PHONY: default build clean install dist build-all upload

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

dist: build-all upload

build-all: clean build
	mkdir -p build/dist
	GOOS=darwin GOARCH=amd64 go build -o build/dist/$(NAME)-macos $(MAIN_SRC)
	GOOS=windows GOARCH=amd64 go build -o build/dist/$(NAME)-windows $(MAIN_SRC)
	GOOS=linux GOARCH=amd64 go build -o build/dist/$(NAME)-linux64 $(MAIN_SRC)
	GOOS=linux GOARCH=arm GOARM=7 go build -o build/dist/$(NAME)-linuxarm7 $(MAIN_SRC)
	GOOS=linux GOARCH=arm GOARM=6 go build -o build/dist/$(NAME)-linuxarm6 $(MAIN_SRC)

upload:
	aws s3 cp build/dist/$(NAME)-macos s3://awssecret2env/master/$(NAME)-macos --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-windows s3://awssecret2env/master/$(NAME)-windows --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-linux64 s3://awssecret2env/master/$(NAME)-linux64 --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-linuxarm7 s3://awssecret2env/master/$(NAME)-linuxarm7 --content-type application/octet-stream
	aws s3 cp build/dist/$(NAME)-linuxarm6 s3://awssecret2env/master/$(NAME)-linuxarm6 --content-type application/octet-stream
