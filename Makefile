export GO15VENDOREXPERIMENT=1

GLIDE := $(GOPATH)/bin/glide
GLIDE_YAML := glide.yaml
GLIDE_LOCK := glide.lock
GLIDE_SWAGGER_YAML := glide-swagger.yaml
GO_GET := .go.get
SWAGGER_SRC := vendor/github.com/go-swagger/go-swagger/cmd/swagger/swagger.go
SWAGGER := $(CURDIR)/swagger
SPEC_FILE := swagger-spec/monorail.yml
CLIENT := client
MODELS := models
CLIENT_BINDINGS := $(CLIENT)/monorail_client.go
GOCOV := $(GOPATH)/bin/gocov
GOVERALLS := $(GOPATH)/bin/goveralls
COVER := $(GOPATH)/bin/cover
COVER_OUT := cover.out

all: install

deps: $(CLIENT_BINDINGS) $(GO_GET)

$(GO_GET): | $(GLIDE_LOCK)
	go get -d $$($(GLIDE) nv) && touch $@

$(GLIDE_LOCK): $(GLIDE_YAML) | $(GLIDE) $(CLIENT_BINDINGS)
	$(GLIDE) --home $(HOME) up

$(GLIDE):
	go get -u github.com/Masterminds/glide

$(SWAGGER): $(SWAGGER_SRC)
	cd $(dir $(<D)) && go build -o $@ ./swagger && cd $(@D)

$(SWAGGER_SRC): $(GLIDE_SWAGGER_YAML) | $(GLIDE)
	$(GLIDE) --home $(HOME) up --quick --use-gopath --file $<

$(CLIENT_BINDINGS): $(SWAGGER) $(SPEC_FILE)
	$(SWAGGER) generate client -f $(SPEC_FILE)

install: $(CLIENT_BINDINGS) $(GO_GET)
	go install -v $$($(GLIDE) nv)

test: $(COVER_OUT)
$(COVER_OUT):
	go test -v $$($(GLIDE) nv) -cover -coverprofile $@

cover: $(COVER_OUT) | $(GOCOV) $(GOVERALLS) $(COVER)
	$(GOVERALLS) -coverprofile=$<

clean:
	rm -fr $(GLIDE_LOCK) $(GO_GET) $(COVER_OUT) $(CLIENT) $(MODELS)

clobber: clean
	rm -fr vendor $(SWAGGER)

.PHONY: all deps install test cover clean clobber
