//go:build !android

package permisions

import (
	"fmt"
)

func RequestPermision(permName string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func IsPermision(permName string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func Init() {
}
