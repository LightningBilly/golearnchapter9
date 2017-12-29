package main

import (
	"fmt"
	"reflect"
)

type T struct{
	A int
	B string
}

func main() {
	/*
	var x float64 = 3.4
	p:=reflect.ValueOf(&x)

	fmt.Println("type:", p.Type())
	fmt.Println(p.CanSet())

	v:=p.Elem()
	fmt.Println(v.CanSet())
	fmt.Println(v.Addr())
	v.SetFloat(4.1)

	fmt.Println(v.Interface())
	fmt.Println(x)
	*/

	t:=T{203, "mh203"}
	s:=reflect.ValueOf(&t).Elem()
	typeOfT:=s.Type()

	for i := 0; i < s.NumField(); i++ {
		f:=s.Field(i)
		fmt.Printf("%d : %s %s = %v\n", i, 
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
