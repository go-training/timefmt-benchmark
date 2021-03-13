# timefmt-benchmark

## Benchmark Result

Time Format: `%Y-%m-%d`, benchmark command

```sh
go test -v -count 10 -bench . | tee benchmark.txt
```

See [the result](./benchmark.txt)

```sh
$ benchstat benchmark.txt 
name             time/op
StdTimeFormat-4   167ns ± 4%
Imperfectgo-4     135ns ± 7%
Itchyny-4        96.0ns ± 2%
Appleboy-4       74.5ns ± 2%
Fastly-4          171ns ± 2%
Jehiah-4          495ns ± 2%
Lestrrat-4        213ns ± 1%
Tebeka-4         1.11µs ± 4%

name             alloc/op
StdTimeFormat-4   16.0B ± 0%
Imperfectgo-4     0.00B     
Itchyny-4         0.00B     
Appleboy-4        0.00B     
Fastly-4          16.0B ± 0%
Jehiah-4          48.0B ± 0%
Lestrrat-4        80.0B ± 0%
Tebeka-4          64.0B ± 0%

name             allocs/op
StdTimeFormat-4    1.00 ± 0%
Imperfectgo-4      0.00     
Itchyny-4          0.00     
Appleboy-4         0.00     
Fastly-4           1.00 ± 0%
Jehiah-4           6.00 ± 0%
Lestrrat-4         2.00 ± 0%
Tebeka-4           7.00 ± 0%
```
