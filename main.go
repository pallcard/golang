package main

import "fmt"

func main() {
	fmt.Println("vim") //ys iw "   增加“
	fmt.Println("vim") //cs '  "   修改‘为”
	fmt.Println("vim") //ds "      删除“

	s := "abc"
	t := s
	s += "d"
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(&s == &t) //false

	n := 0
	for _, _ = range "hello, 世界" {
		n++
	}
	fmt.Println(n)

}
