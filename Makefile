NAME=chanshare
OS=$(shell uname)

buid: 
	go build -o build/$(NAME)

build-win:
	go build -o build/$(NAME).exe

.PHONY: build build-win
