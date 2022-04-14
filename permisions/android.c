// +build android

#include <android/log.h>
#include <jni.h>
#include <stdbool.h>
#define LOG_FATAL(...) __android_log_print(ANDROID_LOG_FATAL, "Fyne", __VA_ARGS__)

jint PERMISSION_GRANTED ;
jclass ClassPackageManager ;
jfieldID lid_PERMISSION_GRANTED;
jclass ClassContext ;
jmethodID MethodcheckSelfPermission ;
jclass ClassActivity;
jmethodID MethodrequestPermissions;

static jclass find_class(JNIEnv *env, const char *class_name) {
	jclass clazz = (*env)->FindClass(env, class_name);
	if (clazz == NULL) {
		(*env)->ExceptionClear(env);
		LOG_FATAL("cannot find %s", class_name);
		return NULL;
	}
	return clazz;
}

static jmethodID find_method(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
	jmethodID m = (*env)->GetMethodID(env, clazz, name, sig);
	if (m == 0) {
		(*env)->ExceptionClear(env);
		LOG_FATAL("cannot find method %s %s", name, sig);
		return 0;
	}
	return m;
}

static jmethodID find_static_method(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
	jmethodID m = (*env)->GetStaticMethodID(env, clazz, name, sig);
	if (m == 0) {
		(*env)->ExceptionClear(env);
		LOG_FATAL("cannot find method %s %s", name, sig);
		return 0;
	}
	return m;
}
static jstring android_permission_name(JNIEnv* env , const char* perm_name) {
    jclass ClassManifestpermission = (*env)->FindClass(env, "android/Manifest$permission");
    jfieldID lid_PERM = (*env)->GetStaticFieldID(env, ClassManifestpermission, perm_name, "Ljava/lang/String;");
    jstring ls_PERM = (jstring)((*env)->GetStaticObjectField(env, ClassManifestpermission, lid_PERM));
    return ls_PERM;
}

void android_permission_init(JNIEnv* env ){
    PERMISSION_GRANTED = (jint)-1;
    ClassPackageManager = find_class(env, "android/content/pm/PackageManager");
    lid_PERMISSION_GRANTED = (*env)->GetStaticFieldID(env, ClassPackageManager,"PERMISSION_GRANTED", "I");
    PERMISSION_GRANTED = (*env)->GetStaticIntField(env, ClassPackageManager, lid_PERMISSION_GRANTED);
    ClassContext = (*env)->FindClass(env, "android/content/Context");
    MethodcheckSelfPermission = (*env)->GetMethodID(env,  ClassContext, "checkSelfPermission", "(Ljava/lang/String;)I");

     ClassActivity =(*env)->FindClass(env, "android/app/Activity");
    MethodrequestPermissions = (*env)->GetMethodID(env,  ClassActivity, "requestPermissions", "([Ljava/lang/String;I)V");
}

bool android_has_permission(JNIEnv* env, jobject ctx , const char* perm_name){
    bool result = false;
    jstring ls_PERM = android_permission_name(env, perm_name);
    jint int_result = (*env)->CallIntMethod(env, ctx, MethodcheckSelfPermission, ls_PERM);
    result = (int_result == PERMISSION_GRANTED);
    return result;
}

void android_request_permissions(JNIEnv* env, jobject ctx , const char* perm_name){
    jobjectArray perm_array = (*env)->NewObjectArray(env,
          1,
          (*env)->FindClass(env, "java/lang/String"),
          (*env)->NewStringUTF(env, "")
        );
    jstring ls_PERM = android_permission_name(env, perm_name);
    (*env)->SetObjectArrayElement(env, perm_array, 0, ls_PERM);
    jsize len = (*env)->GetArrayLength(env, perm_array);
    if (len == 0){
        __android_log_write(1, "permision", "len is 0");
        return;
    }
    (*env)->CallVoidMethod(env, ctx, MethodrequestPermissions, perm_array, 0);
}