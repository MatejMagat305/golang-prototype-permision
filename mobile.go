package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"github.com/MatejMagat305/golang-prototype-permision/permisions"
)

func bluetooth() {
	fmt.Println("bluetooth")
	err := app.RunOnJVM(func(vm, env, ctx uintptr) error {
		permisions.Init(vm, env, ctx)
		return nil
	})
	if err != nil {
		fmt.Println("bluetooth ", err)
		return
	}
	err = app.RunOnJVM(func(vm, env, ctx uintptr) error {

		b, err0 := permisions.CheckPermission(vm, env, ctx, "WRITE_EXTERNAL_STORAGE")
		if err0 != nil {
			return err0
		}
		fmt.Println("bluetooth is granted? ", b)
		if !b {
			_, err0 = permisions.CheckOrRequestPermissionSuscess(vm, env, ctx, "WRITE_EXTERNAL_STORAGE")
			if err0 != nil {
				return err0
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("bluetooth ", err)
		return
	}
}
