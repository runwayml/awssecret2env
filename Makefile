NAME := awssecret2env
MAIN_SRC := cmd/$(NAME)/main.go

.PHONY: default build clean install build-all

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

build-all: clean default
	mkdir -p build/package/$(NAME)-macos build/package/$(NAME)-windows build/package/$(NAME)-linux-x64 build/package/$(NAME)-linux-arm7 build/package/$(NAME)-linux-arm6
	GOOS=darwin GOARCH=amd64 go build -o build/package/$(NAME)-macos/$(NAME) $(MAIN_SRC)
	GOOS=windows GOARCH=amd64 go build -o build/package/$(NAME)-windows/$(NAME) $(MAIN_SRC)
	GOOS=linux GOARCH=amd64 go build -o build/package/$(NAME)-linux-x64/$(NAME) $(MAIN_SRC)
	GOOS=linux GOARCH=arm GOARM=7 go build -o build/package/$(NAME)-linux-arm7/$(NAME) $(MAIN_SRC)
	GOOS=linux GOARCH=arm GOARM=6 go build -o build/package/$(NAME)-linux-arm6/$(NAME) $(MAIN_SRC)
	cd build/package/ && \
		tar czf $(NAME)-linux-x64.tar.gz $(NAME)-linux-x64/ && \
		tar czf $(NAME)-linux-arm7.tar.gz $(NAME)-linux-arm7/ && \
		tar czf $(NAME)-linux-arm6.tar.gz $(NAME)-linux-arm6/ && \
		zip -r -9 $(NAME)-macos.zip $(NAME)-macos/ && \
		zip -r -9 $(NAME)-windows.zip $(NAME)-windows/
	rm -rf build/package/$(NAME)-macos build/package/$(NAME)-windows build/package/$(NAME)-linux-x64 build/package/$(NAME)-linux-arm7 build/package/$(NAME)-linux-arm6
