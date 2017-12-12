[Go by example](https://gobyexample.com/reading-files)  
  
```go
package main

import(
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){

	// Only reads the entire file and does not report EOF as an error so != nil
	dat, err := ioutil.ReadFile("/tmp/dat")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))

	f, err := os.Open("/tmp/dat") //* Works even if err as already been declared &/| initialized
	if err != nil{
		panic(err)
	}

	//* Close the file when scope ends
	defer f.Close()
	
	b1 := make([]byte, 5)
	// Reads for the len(b1) into b1, and returns n1 = number of bytes read
	n1, err := f.Read(b1) 	// https://godoc.org/os#File.Read
	if err != nil{
		panic(err)
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0) // https://godoc.org/os#File.Seek  -  offset at 6 | whence = from where and returns the current offset at the end
	if err != nil {
		panic(err)
	}
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0) // maybe a bit like a cursor
	if err != nil{
		panic(err)
	}
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err := f.Seek(0, 0) // like going back to the begining of the file
	if err != nil{
		panic(err)
	}

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)	// returns the next 5 bytes without advancing the reader, if less than 5 -> err
	if err != nil {
		panic(err)
	}
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close has been defered before
}
```
----------------------------------

***Yes, i did rewrite everything myself, not a copy/paste ...*** :sleepy: