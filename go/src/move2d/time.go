package move2d


import (
	"fmt"
	"math/rand"
	"io"
)

func main() {
	randy := rand.Intn(2)
	fmt.Println(randy)

	var r io.Reader
	r.Read([]byte("1123"))
}
