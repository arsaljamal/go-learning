package main

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	// Implementation details
}

func (f *File) Read(p []byte) (n int, err error) {
	// Read implementation
	return 0, nil
}
func (f *File) Write(p []byte) (n int, err error) {
	// Write implementation
	return 0, nil
}

// The File type now satisfies the ReadWriter interface
var rw ReadWriter = new(File)

type A interface {
	Foo()
}
type B interface {
	A
	Foo() //any type implementing interface B must provide 
		  //an implementation for the Foo() method, which will satisfy both A and B interfaces.
}
