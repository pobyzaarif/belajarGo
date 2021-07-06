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
pkg: belajar/concurrency
cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
BenchmarkSumAlgA_100-4         	25547047	     45.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumAlgA_1000-4        	 3006142	     375.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumAlgA_100000000-4   	      30	  36770675 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumAlgB_100-4         	  389761	      3150 ns/op	      56 B/op	       4 allocs/op
BenchmarkSumAlgB_1000-4        	  252988	      4049 ns/op	      56 B/op	       4 allocs/op
BenchmarkSumAlgB_100000000-4   	      60	  21598352 ns/op	      56 B/op	       4 allocs/op
BenchmarkSumAlgC_100-4         	  573788	      2078 ns/op	      96 B/op	       1 allocs/op
BenchmarkSumAlgC_1000-4        	  396218	      2842 ns/op	      96 B/op	       1 allocs/op
BenchmarkSumAlgC_100000000-4   	      58	  22051159 ns/op	      96 B/op	       1 allocs/op
PASS
ok  	belajar/concurrency	12.364s
```