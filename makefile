VERSION ?= 0.0.1

vendoring:
	go get -u github.com/kardianos/govendor
	govendor sync

build: vendoring
	env GOOS=linux GOARCH=amd64 go build .
