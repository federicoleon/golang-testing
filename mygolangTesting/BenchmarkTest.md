# Benchmark Test

Bubble Sort is the worst possible type of Sort Algorithm.
Compare to [Native Sort](api/services/sort_service_test.go#L60)

Using the following Benchmark Tests

```sh
$ go test -v -run="none" -bench=. -benchtime="3s"
goos: darwin
goarch: amd64
pkg: gotestingintegmygolang-testing/api/services
BenchmarkBubbleSort-8           339131985               10.6 ns/op
BenchmarkSort-8                 35946357                97.6 ns/op
BenchmarkBubbleSortS-8            568296              6406 ns/op
BenchmarkBSortLimit-8              46938             76386 ns/op
BenchmarkSortS-8                   46779             77380 ns/op
PASS
```

```sh
$ go test -v -run="none" -bench=. -benchtime="3s"
goos: darwin
goarch: amd64
pkg: gotestingintegmygolang-testing/api/services
BenchmarkBubbleSort-8           366404206               10.3 ns/op
BenchmarkSort-8                 29976718               102 ns/op
BenchmarkBubbleSortS-8            562191              5897 ns/op
BenchmarkBSortLimit-8              49495             73565 ns/op
BenchmarkSortS-8                   48375             71206 ns/op
PASS
ok      gotestingintegmygolang-testing/api/services    22.851s
```

 go test -v -run="none" -bench=. -benchtime="3s"
goos: darwin
goarch: amd64
pkg: gotestingintegmygolang-testing/api/services
BenchmarkBubbleSort-8           366011378                9.73 ns/op
BenchmarkSort-8                 37720219                91.4 ns/op
BenchmarkBubbleSort10X3-8         587432              5848 ns/op
BenchmarkBSortLimitGT10X3-8        51031             69848 ns/op
BenchmarkSort10X3-8                49868             71134 ns/op
PASS
ok      gotestingintegmygolang-testing/api/services    20.166s