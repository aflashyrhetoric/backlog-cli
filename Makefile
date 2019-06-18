.PHONY: init build build-install deps

default: build-install

# FIXME: Later, this should be achievable via blg not Makefile
init: 
	cp backlog-cli.example.yaml $$HOME/backlog-config.yaml

build: 
	go build *.go

build-install:
	go build -o /usr/local/bin/blg *.go

deps:
	dep ensure -v