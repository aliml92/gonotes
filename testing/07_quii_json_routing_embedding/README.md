## Notes

### *Embedding*
> Only interfaces can be embedded within interfaces

Ex. `io.ReadWriter` interface embeds `io.Reader` and `io.Writer` interfaces
```go
type ReadWriter interface {
	Reader
	Writer
}
// Implementations must not retain p.
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Implementations must not retain p.
type Writer interface {
	Write(p []byte) (n int, err error)
}
```
*gopher*: if there is an interface, it is there for the people to implement it.  
*me*: can you give some examples of such implementation, bro?  
*gopher*: sure, for example, `bufio` package has such implementation.  
*me*: where?  
*gopher*: here ⬇️  
```go
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
	*Reader
	*Writer
}
func (b *Reader) Read(p []byte) (n int, err error) {
    // implementation
}
func (b *Writer) Write(p []byte) (nn int, err error) {
    // implementation
}
```
*me*: is there more to learn about embedding?  
*gopher*: visit [embedding in go](https://octo.vmware.com/golang-embedding/)

