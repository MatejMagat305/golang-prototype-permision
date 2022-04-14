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
)

var (
	wasInit = false
)

func CheckOrRequestPermissionSuscess(vm, env, ctx uintptr, permName string) (bool, error) {
	fmt.Println("CheckOrRequestPermissionSuscess")
	if !wasInit {
		return false, fmt.Errorf("it was not initialization")
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
	return bool(C.android_has_permission(envP, ctxP, cPermName)), nil
}

func CheckPermission(vm, env, ctx uintptr, permName string) (bool, error) {
	fmt.Println("CheckPermission")
	if !wasInit {
		return false, fmt.Errorf("it was not initialization")
	}
	cPermName := C.CString(permName)
	defer C.free(unsafe.Pointer(cPermName))
	envP := (*C.JNIEnv)(unsafe.Pointer(env))
	ctxP := (C.jobject)(unsafe.Pointer(ctx))
	return bool(C.android_has_permission(envP, ctxP, cPermName)), nil
}

func Init(vm, env, ctx uintptr) {
	fmt.Println("Init")
	wasInit = true
	envP := (*C.JNIEnv)(unsafe.Pointer(env))
	C.android_permission_init(envP)
}
