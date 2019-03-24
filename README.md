# co2go

Benchmark CO2 emissions for Go programs.

Use as a Go drop-in replacement.

```
~/s/g/co2go ./co2go test -bench . --benchmem
goos: darwin
goarch: amd64
BenchmarkFoo-8    	  300000	      5565 ns/op	       0 B/op	       0 allocs/op	0.000001352088889 g CO2/op
BenchmarkFoo2-8   	  300000	      5530 ns/op	       0 B/op	       0 allocs/op	0.000001343585185 g CO2/op
PASS
ok  	_/Users/gustav/src/go-paris	3.456s
```
