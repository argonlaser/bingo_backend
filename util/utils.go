package util

import (
    "reflect"
    "runtime"
)

// Use GetFunctionName(func)
func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}