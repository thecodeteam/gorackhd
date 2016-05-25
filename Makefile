export GO15VENDOREXPERIMENT=1

GLIDE := $(GOPATH)/bin/glide
GLIDE_LOCK := glide.lock
SWAGGER_PKG := github.com/go-swagger/go-swagger
SWAGGER_CMD := $(SWAGGER_PKG)/cmd/swagger/swagger.go
VENDORED_SWAGGER_CMD := vendor/$(SWAGGER_CMD)
SPEC_FILE := swagger-spec/monorail.yml
CLIENT := client
MODELS := models
CLIENT_BINDINGS := $(CLIENT)/monorail_client.go

all: install

$(GLIDE):
	go get -u github.com/Masterminds/glide

swagger_deps: $(GLIDE_LOCK) | $(GLIDE)
	$(GLIDE) --home $(HOME) install

$(CLIENT_BINDINGS): swagger_deps $(SPEC_FILE)
	go run $(VENDORED_SWAGGER_CMD) generate client -f $(SPEC_FILE)

install: $(CLIENT_BINDINGS)
	go install -v $$($(GLIDE) nv)

test:
	go test -v

clean:
	rm -fr $(COVER_OUT) $(CLIENT) $(MODELS)

clobber: clean
	rm -fr vendor

.PHONY: all swagger_deps install test \
	clean clobber
