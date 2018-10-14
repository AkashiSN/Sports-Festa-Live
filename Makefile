GCFLAGS := -gcflags="-trimpath=${GOPATH}"

build:
	go build ${GCFLAGS} ${LDFLAGS}

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif

.PHONY: deps
deps:
	dep ensure -v
