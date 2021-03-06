[Go by example](https://gobyexample.com/writing-files)  
  
```go
import(
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	//* #1
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1")
	// fmt.Println(d1) //=> [104 101 108 108 111 10 103 111 10]
	check(err)

	//* #2
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	{
		// Commiting changes
		defer f.Sync()

		d2 := []byte{115, 111, 109, 101, 10}
		n2, err := f.Write(d2) // on the file 'f' we Write d2 | Returns the number of bytes written
		check(err)
		fmt.Printf("wrote %d bytes\n", n2)

		n3, err := f.WriteString("writes\n")
		fmt.Printf("wrote %d bytes\n", n3)
	}

	//---
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	
}
```