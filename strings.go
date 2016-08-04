package main

import (
"fmt"
"strings"
)

func main() {

str := "1,2,3,4"

split := strings.Split(str,",")

fmt.Printf("%d",split[0])
}
