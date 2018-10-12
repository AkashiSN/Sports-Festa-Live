GCFLAGS := -gcflags="-trimpath=${GOPATH}"

build:
	go build ${GCFLAGS} ${LDFLAGS}

.PHONY: deps
deps: dep
	dep ensure -v
