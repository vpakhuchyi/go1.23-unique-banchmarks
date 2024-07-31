# go1.23-unique-banchmarks

This repository contains the benchmarks for the "unique" package that is going to be added in Go 1.23.

All the tests are run on the following machine:
- goos: darwin
- goarch: arm64
- cpu: Apple M2 Pro

Go version: go1.23rc2

Link to the "unique" package: https://pkg.go.dev/unique@master


## Benchmarks

```text
BenchmarkString/regular-12   	         800237403	   1.483  ns/op
BenchmarkString/unique-12    	         1000000000	   0.2944 ns/op

BenchmarkSimpleStruct/regular-12         450721364	   2.640  ns/op
BenchmarkSimpleStruct/unique-12          1000000000	   0.2963 ns/op

BenchmarkComplexStruct/regular-12        162752110	   7.365  ns/op
BenchmarkComplexStruct/unique-12         1000000000	   0.2947 ns/op
```