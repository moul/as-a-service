GOENV ?=	GO15VENDOREXPERIMENT=1
GODEP ?=	$(GOENV) godep
GO ?=		$(GOENV) go
SOURCES :=	$(shell find . -name "*.go")

all: build

.PHONY: build
build: moul-as-a-service

.PHONY: test
test:
	$(GODEP) restore
	$(GO) get -t .
	$(GO) test -v .

.PHONY: godep-save
godep-save:
	$(GODEP) save $(shell go list ./... | grep -v /vendor/)

.PHONY: cover
	rm -f profile.out
	$(GO) test -covermode=count -coverpkg=. -coverprofile=profile.out

.PHONY: convey
convey:
	$(GO) get github.com/smartysteets/goconvey
	goconvey -cover -port=9032 -workDir="$(shell realpath .)" -depth=-1

.PHONY: clean
clean:
	rm -rf moul-as-a-service

moul-as-a-service: $(SOURCES)
	$(GO) build -o $@ ./cmd/$@

.PHONY: goapp_serve
goapp_serve:
	goapp serve ./appspot/app.yaml


.PHONY: goapp_deploy
goapp_deploy:
	goapp deploy -application moul-as-a-service ./appspot/app.yaml
