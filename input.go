package main

import "fmt"

func main() {
	fmt.Println("Enter Your && Your Partner Name:")
	var (
		me string
		you string
	)
	fmt.Scanf("%s%s",&me,&you)
	fmt.Printf("%s Love %s",me,you)

}
