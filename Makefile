.PHONY: build build-install moveToBin all

PWD = $(shell pwd)

default: build-install

all: build-install

# FIXME: Later, this should be achievable via blg not Makefile
init: 
	cp backlog-cli.example.yaml $$HOME/backlog-config.yaml

build: 
	go build *.go

build-install:
	go build -o /usr/local/bin/blg *.go