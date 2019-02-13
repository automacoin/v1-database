# v1-database

Database component for the v1.0 architecture of AutomaCoin - https://www.automacoin.com/

## Prerequisites

* [Go Version 1.11+](https://golang.org/dl/)

## Dependencies

[Modules](https://github.com/golang/go/wiki/Modules) are implemented
after go v1.11. No need to `go get` anymore!

## Compile

	make database

## Running the program

After building, you can

	./build/v1-database <command-line-options>

## Testing

To run your tests, do

	make test

Or you can do `go test`, or you can have the `ginkgo` executable,

	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega/...

This fetches ginkgo and installs the ginkgo executable under
`$GOPATH/bin` – you’ll want that on your `$PATH`.

and then

	ginkgo

## License

MIT

## Contributors

Just fill an issue here!