# This is how we want to name the binary output
BINARY=primestring

# Project name
PROJECT="primestring"

# These are the values we want to pass for Version and BuildTime
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

# build directory
BUILD_DIR=${GOBIN}

# command path
CMD=cmd

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

all:
	go build $(LDFLAGS) -o primestring primestring.go 
#	go build $(LDFLAGS) -o primestringtwo primestringTwo.go 
	go build $(LDFLAGS) -o primestringhex primestringHex.go 

clean:
	rm -r primestring primestringtwo primestringhex
