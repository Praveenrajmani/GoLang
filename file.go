package main

import (
//"fmt"
"os"
"log"
"fmt"
"io/ioutil"
)

func main() {

newfile, err := os.Create("File.txt")

if err != nil {
log.Fatal(err)
}
newfile.WriteString("Hello")
newfile.Close()
str, err := ioutil.ReadFile("File.txt")

Strread := string(str)

fmt.Println(Strread)


}
