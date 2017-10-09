PROJECT_ROOT=$(shell pwd)
GOBASE="$(PROJECT_ROOT)/gae/gopath:$(PROJECT_ROOT)/vendor"

deploy:
	GOPATH=$(GOBASE) gcloud app deploy ./gae/api/api.yaml
.PHONY: deploy

run:
	GOPATH=$(GOBASE) goapp serve ./gae/api/api.yaml
.PHONY: run

gopath:
	@echo $(GOBASE)
.PHONY: gopath

add-index:
	GOPATH=$(GOBASE) gcloud datastore create-indexes ./gae/api/index.yaml
.PHONY: add-index

remove-index:
	GOPATH=$(GOBASE) gcloud datastore cleanup-indexes ./gae/api/index.yaml
.PHONY: remove-index
