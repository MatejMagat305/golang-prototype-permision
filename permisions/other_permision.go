//go:build !android

package permisions

import (
	"fmt"
)

func CheckOrRequestPermissionSuscess(vm, env, ctx uintptr, permName string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func CheckPermission(vm, env, ctx uintptr, permName string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func Init(vm, env, ctx uintptr) {
}
