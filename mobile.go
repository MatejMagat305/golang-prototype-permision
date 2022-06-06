package main

import (
	"fmt"
	"github.com/MatejMagat305/golang-prototype-permision/permisions"
)

func bluetooth() {
	fmt.Println("bluetooth")
	permisions.Init()
	b, err0 := permisions.IsPermision("WRITE_EXTERNAL_STORAGE")
	if err0 != nil {
		return
	}
	fmt.Println("bluetooth is granted? ", b)
	if !b {
		b, err0 = permisions.RequestPermision("WRITE_EXTERNAL_STORAGE")
		if err0 != nil {
			return
		}
		fmt.Println("bluetooth is granted? ", b)
	}
}