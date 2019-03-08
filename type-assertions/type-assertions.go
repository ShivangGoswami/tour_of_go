package main

import "fmt"

func main(){
	var i interface{}="hello"

	s:=i.(string)
	fmt.Println(s)

	s,ok:=i.(string)
	fmt.Println(s,ok)

	i=64.0
	f:=i.(float64)
	fmt.Println(f)
}