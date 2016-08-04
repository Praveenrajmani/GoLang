package main
import "fmt"
import "time"

func main() {
    t0 := time.Now()
    
    resulting := 0
    for resulting<5 {
    fmt.Println("Hi \n")

    //time.Sleep(3000 * time.Millisecond)
    t1 := time.Now()
    resulting = int(t1.Sub(t0)/(time.Millisecond*1000))
    fmt.Println("\n TIME NOW:  ",resulting)
   // if t1.Sub(t0) > 5.0	{
    //fmt.Printf("%#v",t1.Sub(t0)/(time.Millisecond*1000))
    //}
}
}
