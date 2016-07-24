package main
import "fmt"

type mystruct struct {
  length,breadth int
}

func main() {
    r := mystruct{1,3}
    p := mystruct{}
    addr := new(mystruct)
    p.length = 1
    p.breadth = 2
    fmt.Println("r is ",r)
    fmt.Println("p is ",p)
    addr.length = 32
    addr.breadth = 23
    fmt.Println("addr is ",*(addr))
    fmt.Println("breadth is ",addr.breadth)

}
