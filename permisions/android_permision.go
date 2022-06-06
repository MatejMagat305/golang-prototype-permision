//go:build android
// +build android

package permisions

/*
#cgo LDFLAGS: -landroid -llog
#include <jni.h>
#include <stdbool.h>
#include <stdlib.h>
void android_permission_init(JNIEnv* env);
bool android_has_permission(JNIEnv* env, jobject ctx, const char* perm_name);
void android_request_permissions(JNIEnv* env, jobject ctx, const char* perm_name);
*/
import "C"
import (
	"fmt"
	"unsafe"
	_ "unsafe"
)

var (
	wasInit = false
)

//go:linkname runOnJVM fyne.io/fyne/v2/internal/driver/mobile/mobileinit.RunOnJVM
func runOnJVM(fn func(vm, env, ctx uintptr) error) error

func isPermision(permName string) (b bool, e error) {
	runOnJVM(func(vm, env, ctx uintptr) error {
		fmt.Println("CheckPermission")
		if !wasInit {
			b, e = false, fmt.Errorf("it was not initialization")
			return e
		}
		cPermName := C.CString(permName)
		defer C.free(unsafe.Pointer(cPermName))
		envP := (*C.JNIEnv)(unsafe.Pointer(env))
		ctxP := (C.jobject)(unsafe.Pointer(ctx))
		b, e = bool(C.android_has_permission(envP, ctxP, cPermName)), nil
		return nil
	})
}

func initEnv() {
	runOnJVM(func(vm, env, ctx uintptr) error {
		fmt.Println("Init")
		wasInit = true
		envP := (*C.JNIEnv)(unsafe.Pointer(env))
		C.android_permission_init(envP)
	})
}

func RequestPermision(permName string) (b bool, e error) {
	runOnJVM(func(vm, env, ctx uintptr) error {
		fmt.Println("CheckOrRequestPermissionSuscess")
		if !wasInit {
			b, e = false, fmt.Errorf("it was not initialization")
			return e
		}
		cPermName := C.CString(permName)
		defer C.free(unsafe.Pointer(cPermName))
		envP := (*C.JNIEnv)(unsafe.Pointer(env))
		ctxP := (C.jobject)(unsafe.Pointer(ctx))
		has := bool(C.android_has_permission(envP, ctxP, cPermName))
		if !has {
			fmt.Println("CheckOrRequestPermissionSuscess request")
			C.android_request_permissions(envP, ctxP, cPermName)
		}
		b, e = bool(C.android_has_permission(envP, ctxP, cPermName)), nil
		return nil
	})
}
