# Wring

Package `wring` provides a facility to minify HTML code.

## Install 

To install, use `go get`:

```bash
$ go get -u github.com/elpinal/wring
```

## Examples

```go
err := wring.HTML(strings.NewReader("<p>    Go    </p>"), os.Stdout)
```

## Contribution

1. Fork ([https://github.com/elpinal/wring/fork](https://github.com/elpinal/wring/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[elpinal](https://github.com/elpinal)
