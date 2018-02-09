package main

import (
	"myfunc"
	"fmt"
)

func main(){
	myfunc.SayHello()//
	fmt.Println(myfunc.File)
	const freezingF, boilingF =32.0,212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, myfunc.FToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, myfunc.FToC(boilingF))

	p:=myfunc.NewInt()
	q:=new(int)
	fmt.Println(p == q)

}
