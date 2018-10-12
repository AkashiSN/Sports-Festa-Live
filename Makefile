GCFLAGS := -gcflags="-trimpath=${GOPATH}"

build:
	go build ${GCFLAGS} ${LDFLAGS}