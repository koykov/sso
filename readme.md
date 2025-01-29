# SSO

Small String Optimization (SSO) is a technique used by `std::string` implementations in C++ to store small strings
directly within the string object itself, rather than allocating memory on the heap. This optimization leverages the
fact that many strings in typical applications are relatively short, and thus can be stored within the internal buffer
of the `std::string` object, avoiding dynamic memory allocation.

### How it works

Any string in Go with non-zero length typically consists of two parts - header structure and space in heap where the
actual data stored. Let's take a closer look at the header structure:
```go
type StringHeaader struct {
    Data uintptr
    Len  int
}
```
In `amd64` architecture it takes up 16 bytes in memory (8 bytes for `Data` and 8 bytes for `Len`). The main idea is to
use that space for strings less than 16 bytes to avoid allocations in heap - similar to `std::string` in C++.

In fact this SSO implementation can store without allocations only 15 bytes (one byte is a header that stores length of
optimized string and special flag is string optimized or not).

Note: in `x86` architecture SSO string may contain 7 bytes without allocations due to smaller size of `StringHeader` structure.

### Usage

```go
import "github.com/koykov/sso"

var s sso.String                    // contains 0 bytes of data; 15 bytes left
s.AssignString("hello")             // contains 5 bytes of data; 10 bytes left
s.AppendString(" world!")           // contains 12 bytes of data; 3 bytes left
println(s.String)                   // "hello world!"
s.AppendString("some long string")  // final size exceeds the limit of 15 bytes, so heap allocation occurs and `s` became regular Go string
println(s.String)                   // "hello world!some long string"
s.Reset()                           // `s` became SSO string and ready to store data up to 15 bytes, previously allocated data became a garbage and GC will care about it
```

### Pros and cons

#### Pros:
* Reduced Heap Allocations: Since many strings are short, SSO can avoid heap allocations and deallocations for a
significant portion of string operations, leading to faster execution times.
* Cache Efficiency: Storing small strings within the object can improve cache locality, as accessing the string data
involves fewer memory accesses.
* Lower Overhead: By avoiding dynamic memory allocation for small strings, SSO reduces the overhead associated with
memory management.
* Less Need for Manual Optimization: Developers can benefit from performance improvements without having to manually
optimize string handling for small strings.

#### Cons:
* Limited Capacity: SSO is only effective for small strings that fit within the internal buffer of the string object.
* Manual handling: Go doesn't support operators overloading, thus special methods must be used for assigning/concatenation of data.
