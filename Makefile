GOENV ?=	GO15VENDOREXPERIMENT=1
GODEP ?=	$(GOENV) godep
GO ?=		$(GOENV) go
SOURCES :=	$(shell find . -name "*.go")
PORT ?=		8000


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
	$(GOENV) goconvey -cover -port=9032 -workDir="$(shell realpath .)" -depth=-1

.PHONY: clean
clean:
	rm -rf moul-as-a-service

moul-as-a-service: $(SOURCES)
	$(GO) build -o $@ ./cmd/$@

.PHONY: goapp_serve
goapp_serve:
	$(GOENV) goapp serve ./appspot/app.yaml


.PHONY: goapp_deploy
goapp_deploy:
	$(GOENV) goapp deploy -application moul-as-a-service ./appspot/app.yaml

.PHONY: gin
gin:
	$(GO) get ./...
	$(GO) get github.com/codegangsta/gin
	cd ./cmd/moul-as-a-service; $(GOENV) gin --immediate --port=$(PORT) server


.PHONY: heroku_deploy
heroku_deploy:
	#git remote add heroku https://git.heroku.com/moul-showcase.git
	git push heroku master


.PHONY: dokku_deploy
dokku_deploy:
	#git remote add dokku dokku@dokku.m.42.am:moul-showcase
	git push dokku master
