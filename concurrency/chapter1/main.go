package main

import "fmt"

func main() {
	i := 1
	j := 2
	i, j = j, i
	fmt.Println("vim-go")
	fmt.Println(i)
	fmt.Println(j)
}
