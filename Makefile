targets = backend

.PHONY: build
build: $(targets)

$(targets): $(targets).go
	CGO_ENABLED=0 go build -o $@ $@.go

.PHONY: vet
g := ~/go/bin/gocritic
vet:
	go vet
	if [ -x $g ]; then $g check -enable='#style'; fi

.PHONY: test
test: vet
	go test

.PHONY: clean
clean:
	rm -f -- $(targets)
