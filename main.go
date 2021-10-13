package main

import (
	"easy-func/string/stristr"
	"fmt"
)

func main(){
	s, err := stristr.Do("fasfasdf 333", "33" , true)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(s)
}