package main

import "fmt"

func main() {
	arr := []int64{1,2,3,4,5,6,7,8,9,10}
	for _ , value := range arr {
		if(value % 2 ==0) {	
			fmt.Println("Even")
		}else 
		{
			fmt.Println("odd")
		}
	}
}