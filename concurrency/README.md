case :
calculate and sum number from 1 to n...
we don't care about cpu or memory usage, the aim is maximize the speed of proces to get sum number

suggestion :
try compare sequential, parallel or concurrency

result :
$ go test -bench=. -benchmem
```
goos: darwin
goarch: amd64
pkg: github.com/pobyzaarif/belajarGo/concurrency
cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
BenchmarkSumAlgA_100-4         	19195593	     65.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumAlgA_1000-4        	 2532300	     414.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumAlgA_100000000-4   	      30	  40843531 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumAlgB_100-4         	  307144	      3616 ns/op	     248 B/op	       8 allocs/op
BenchmarkSumAlgB_1000-4        	  238645	      4314 ns/op	     248 B/op	       8 allocs/op
BenchmarkSumAlgB_100000000-4   	      58	  23121940 ns/op	     248 B/op	       8 allocs/op
BenchmarkSumAlgC_100-4         	  499446	      2409 ns/op	     192 B/op	       3 allocs/op
BenchmarkSumAlgC_1000-4        	  368048	      3048 ns/op	     192 B/op	       3 allocs/op
BenchmarkSumAlgC_100000000-4   	      54	  21770005 ns/op	     192 B/op	       3 allocs/op
PASS
ok  	github.com/pobyzaarif/belajarGo/concurrency	11.860s
```