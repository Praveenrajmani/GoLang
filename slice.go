package main

import "fmt"

func main() {

	 
	Names := []string {"Praveen","Preetha","Mani","Bharathi","Vadivu","Tydy"}
	PresAge := make(map[string] int)
	PresAge[Names[0]] = 21
	PresAge[Names[1]] = 18
	PresAge[Names[2]] = 50
	PresAge[Names[3]] = 40
	PresAge[Names[4]] = 60
	PresAge[Names[5]] = 2
	for _, value := range Names {
		fmt.Printf("PresAge[%s] = %d \n",value,PresAge[value])
	}
}

