package main

import (
	"fmt"
	"io"
	"os"
)



func main() {
	f, err := os.Open(os.Args[1])
 	if(err != nil) {
		fmt.Println("Error" , err)
		os.Exit(1)
	}
	fmt.Println(f)
	io.Copy(os.Stdout, f)
}


func readFile(fileName string)   (s string){
	bs,err := os.ReadFile(fileName)
	if err != nil  {
		fmt.Println("Error",err)
		os.Exit(1)
	}

	s = string(bs)
	return  s
}