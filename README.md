# Testing in GOLANG

![](https://womanonrails.com/images/tdd-basics/tdd.gif)

## Usage

### Only test

```bash
go test -v
```

### Coverage

```bash
go test -cover
```
or
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Benchmark

```bash
go test -bench=.
go test -bench=ShowHello
```

### Production

```bash
go test && go run .
```


## Refrence

* https://blog.alexellis.io/golang-writing-unit-tests/
* https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package