PROJECT_ROOT=$(shell pwd)
GOPATHS="$(PROJECT_ROOT)/gae/gopath:$(PROJECT_ROOT)/vendor"

deploy:
	GOPATH=$(GOPATHS) gcloud app deploy ./gae/api/api.yaml
.PHONY: deploy

run:
	GOPATH=$(GOPATHS) goapp serve ./gae/api/api.yaml
.PHONY: run

gopath:
	@echo $(GOPATHS)
.PHONY: gopath

add-index:
	GOPATH=$(GOPATHS) gcloud datastore create-indexes ./gae/api/index.yaml
.PHONY: add-index

remove-index:
	GOPATH=$(GOPATHS) gcloud datastore cleanup-indexes ./gae/api/index.yaml
.PHONY: remove-index
