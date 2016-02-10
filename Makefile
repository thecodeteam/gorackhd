# recursively list the contents of a directory
RLSDIR = $(wildcard $1) $(foreach d,$(wildcard $1*),$(call RLSDIR,$d/))
PKG_DIRS_IGNORE_PATTS := ./vendor% ./swagger-spec%
PKG_DIRS := $(filter-out $(PKG_DIRS_IGNORE_PATTS),$(sort $(dir $(call RLSDIR,./))))

all: install

deps:
	go get -d $(PKG_DIRS)
	go get -v github.com/axw/gocov/gocov
	go get -v github.com/mattn/goveralls
	go get -v golang.org/x/tools/cmd/cover
	go get -v github.com/go-swagger/go-swagger

client/monorail_client.go: swagger-spec/monorail.yml
	rm -fr client models && \
		swagger generate client -f swagger-spec/monorail.yml

install: client/monorail_client.go
	go install -v $(PKG_DIRS)

test: cover.out
cover.out:
	go test -v $(PKG_DIRS) -cover -coverprofile cover.out

cover: cover.out
	goveralls -coverprofile=cover.out

.PHONY: all deps install test cover
