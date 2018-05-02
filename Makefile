.PHONY: build 

all:
	@echo "make <build|etc...>"

print-%: ; @echo $*=$($*)


##
## Building
##
dist: 
	@go build -i

build:
	go build -i

test:
	go test
