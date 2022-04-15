//go:build !android

package permisions

import (
	"fmt"
)

func RequestPermision(vm, env, ctx uintptr, permName string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func IsPermision(vm, env, ctx uintptr, permName string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func Init(vm, env, ctx uintptr) {
}
