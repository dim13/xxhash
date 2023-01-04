# xxhash

Go implementation of [xxHash][1] (XXH32, XXH64).

## benchmark (Air/M1)

	BenchmarkXXH32-8   	 7964274	       150.5 ns/op	       0 B/op	       0 allocs/op
	BenchmarkXXH64-8   	11149471	       104.4 ns/op	       0 B/op	       0 allocs/op

[1]: https://cyan4973.github.io/xxHash/
