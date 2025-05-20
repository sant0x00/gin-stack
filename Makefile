.PHONY: build-* run-*

build-flat-controller-binding:
	go build -o bin/flat-controller-binding examples/flat-controller-binding/*.go

build-modular-controllers:
	go build -o bin/modular-controllers examples/modular-controllers/*.go

run-flat-controller-binding:
	make build-flat-controller-binding
	./bin/flat-controller-binding

run-modular-controllers:
	make build-modular-controllers
	./bin/modular-controllers