Contributing
============
[Where to start contributing](https://github.com/rackerlabs/maestro/issues?utf8=✓&q=is%3Aissue%20is%3Aopen%20no%3Aassignee%20label%3A"help%20wanted")

### Development Requirements
- Go version 1.9 or later - [Installation docs](https://golang.org/doc/install)
- [Godep](https://github.com/Masterminds/glide)
```shell
go get github.com/Masterminds/glide
```
- [Golint](https://github.com/golang/lint)
```shell
go get -u github.com/golang/lint/golint
```
- Make - This isn't required but can be helpful

### Build
You must have a valid [GOPATH](https://golang.org/doc/code.html#GOPATH) setup in order to fetch the the code.

```shell
go get -d github.com/rackerlabs/maestro
```

If Make is installed run `make build` and a new binary should be available and ready
to use in bin.

If make is not available you can also execute the command directly.
```shell
go build -o bin/${COMMAND_NAME}
```

### Testing
If make is installed run all the tests run by CI run `make test`. Otherwise all of the
commands to test the code will need to be run manually. See the [Makefile](Makefile) for
more information.
