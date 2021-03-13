# timefmt-benchmark

## Benchmark Result

```sh
$ go test -v -bench .
goos: darwin
goarch: amd64
pkg: github.com/go-training/timefmt-benchmark
cpu: Intel(R) Core(TM) i5-4278U CPU @ 2.60GHz
BenchmarkStdTimeFormat
BenchmarkStdTimeFormat-4         7130658               165.5 ns/op            16 B/op          1 allocs/op
BenchmarkImperfectgo
BenchmarkImperfectgo-4           8620616               133.8 ns/op             0 B/op          0 allocs/op
BenchmarkItchyny
BenchmarkItchyny-4              12232382                94.26 ns/op            0 B/op          0 allocs/op
BenchmarkFastly
BenchmarkFastly-4                6723265               170.3 ns/op            16 B/op          1 allocs/op
BenchmarkJehiah
BenchmarkJehiah-4                2478412               481.4 ns/op            48 B/op          6 allocs/op
BenchmarkLestrrat
BenchmarkLestrrat-4              5539683               213.6 ns/op            80 B/op          2 allocs/op
BenchmarkTebeka
BenchmarkTebeka-4                1084783              1100 ns/op              64 B/op          7 allocs/op
PASS
ok      github.com/go-training/timefmt-benchmark        10.966s
```
