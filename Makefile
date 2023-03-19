NAME = $(shell basename  ${PWD})

build:
	CGO_ENABLED=0 go build -o bin/${NAME}