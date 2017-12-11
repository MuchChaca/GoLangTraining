[Go by example](https://gobyexample.com/reading-files)
```go
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
		panic(e)
	}
	fmt.Println(string(dat))

	f, err := os.Open("/tmp/dat")
	//* <notOver>
}
```