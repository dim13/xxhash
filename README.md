[![Build Status](https://travis-ci.org/dim13/xxhash.svg?branch=master)](https://travis-ci.org/dim13/xxhash)
[![GoDoc](https://godoc.org/github.com/dim13/xxhash?status.svg)](https://godoc.org/github.com/dim13/xxhash)

# xxhash is a non-cryptographic hash algorithm

This package implements [xxHash](https://cyan4973.github.io/xxHash/) in Go

```
BenchmarkXXHash32-4   	 1000000	      1017 ns/op	     104 B/op	      26 allocs/op
BenchmarkXXHash64-4   	 2000000	       698 ns/op	     104 B/op	      14 allocs/op
```

Tested on 1.3 GHz Intel Core i5
