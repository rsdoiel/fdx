#
# Simple Makefile
#
PROJECT = fdx

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif


build: bin/fdx2txt$(EXT) bin/txt2fdx$(EXT)

bin/fdx2txt$(EXT): fdx.go fromfountain.go cmd/fdx2txt/fdx2txt.go
	go build -o bin/fdx2txt$(EXT) cmd/fdx2txt/fdx2txt.go

bin/txt2fdx$(EXT): fdx.go fromfountain.go cmd/txt2fdx/txt2fdx.go
	go build -o bin/txt2fdx$(EXT) cmd/txt2fdx/txt2fdx.go

man: build
	mkdir -p man/man1
	bin/fdx2txt -generate-manpage | nroff -Tutf8 -man > man/man1/fdx2txt.1
	bin/txt2fdx -generate-manpage | nroff -Tutf8 -man > man/man1/txt2fdx.1
	
test:
	go test

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

website:
	./mk-website.bash

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

publish:
	./mk-website.bash
	./publish.bash

