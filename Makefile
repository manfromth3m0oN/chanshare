NAME=chanshare
OS=$(shell uname)

buid: 
	go build -o build/$(NAME)

build-win:
	go build -o build/$(NAME).exe

run:
	./build/chanshare

.PHONY: build build-win
