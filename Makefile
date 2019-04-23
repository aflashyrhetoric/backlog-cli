.PHONY: build moveToBin all

PWD = $(shell pwd)

default: all

all: build moveToBin

# FIXME: Later, this should be achievable via blg not Makefile
init: 
	cp backlog-cli.example.yaml $$HOME/backlog-config.yaml

build:
	go build *.go

moveToBin:
	mv main $$HOME/blg