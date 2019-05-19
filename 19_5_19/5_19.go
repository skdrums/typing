package main

import (
	"fmt"
)

func main() {
	var A,B,T int
	fmt.Scanf("%d %d %d",&A,&B,&T)
	n:= T/A
	fmt.Print(B*n)

}
 
