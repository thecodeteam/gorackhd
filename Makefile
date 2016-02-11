export GO15VENDOREXPERIMENT=1

all: install

deps:
	go get github.com/Masterminds/glide
	glide --home $(HOME) up
	go get -d $$(glide nv)
	go get -v github.com/axw/gocov/gocov
	go get -v github.com/mattn/goveralls
	go get -v golang.org/x/tools/cmd/cover

vendor/github.com/go-swagger/go-swagger/cmd/swagger/swagger.go:
	glide --home $(HOME) up

swagger: vendor/github.com/go-swagger/go-swagger/cmd/swagger/swagger.go
	cd vendor/github.com/go-swagger/go-swagger/cmd && \
		go build -o $(CURDIR)/swagger ./swagger && \
		cd -

client/monorail_client.go: swagger swagger-spec/monorail.yml
	rm -fr client models && \
		./swagger generate client -f swagger-spec/monorail.yml

install: client/monorail_client.go
	go install -v $$(glide nv)

test: cover.out
cover.out:
	go test -v $$(glide nv) -cover -coverprofile cover.out

cover: cover.out
	goveralls -coverprofile=cover.out

.PHONY: all deps install test cover
