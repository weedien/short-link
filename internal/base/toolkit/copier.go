package toolkit

//
//import (
//	"fmt"
//	"reflect"
//	"unsafe"
//)
//
//// CopyStruct 将一个结构体的字段拷贝到另一个结构体的相同字段上
//// 两个结构体的字段名和类型必须一致
//func CopyStruct(dst, src interface{}) error {
//	srcVal := reflect.ValueOf(src)
//	dstVal := reflect.ValueOf(dst)
//
//	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
//		return fmt.Errorf("src and dst must be pointers")
//	}
//
//	if srcVal.IsNil() || dstVal.IsNil() {
//		return fmt.Errorf("src and dst must not be nil")
//	}
//
//	srcElem := srcVal.Elem()
//	dstElem := dstVal.Elem()
//
//	if !dstElem.CanSet() {
//		return fmt.Errorf("dst must be settable")
//	}
//
//	copyRecursive(dstElem, srcElem)
//	return nil
//}
//
//// DeepCopy performs a deep copy of the given source struct to the destination struct.
//func DeepCopy(dst, src interface{}) error {
//	srcVal := reflect.ValueOf(src)
//	dstVal := reflect.ValueOf(dst)
//
//	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
//		return fmt.Errorf("src and dst must be pointers")
//	}
//
//	if srcVal.IsNil() || dstVal.IsNil() {
//		return fmt.Errorf("src and dst must not be nil")
//	}
//
//	srcElem := srcVal.Elem()
//	dstElem := dstVal.Elem()
//
//	if !dstElem.CanSet() {
//		return fmt.Errorf("dst must be settable")
//	}
//
//	copyRecursive(dstElem, srcElem)
//	return nil
//}
//
//func copyRecursive(dst, src reflect.Value) {
//	switch src.Kind() {
//	case reflect.Ptr:
//		if src.IsNil() {
//			dst.Set(reflect.Zero(dst.Type()))
//		} else {
//			dst.Set(reflect.New(src.Type().Elem()))
//			copyRecursive(dst.Elem(), src.Elem())
//		}
//	case reflect.Struct:
//		for i := 0; i < src.NumField(); i++ {
//			srcField := src.Field(i)
//			dstField := dst.Field(i)
//			if !dstField.CanSet() {
//				dstField = reflect.NewAt(dstField.Type(), unsafe.Pointer(dstField.UnsafeAddr())).Elem()
//			}
//			copyRecursive(dstField, srcField)
//		}
//	case reflect.Slice:
//		if src.IsNil() {
//			dst.Set(reflect.Zero(dst.Type()))
//		} else {
//			dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
//			for i := 0; i < src.Len(); i++ {
//				copyRecursive(dst.Index(i), src.Index(i))
//			}
//		}
//	case reflect.Map:
//		if src.IsNil() {
//			dst.Set(reflect.Zero(dst.Type()))
//		} else {
//			dst.Set(reflect.MakeMap(src.Type()))
//			for _, key := range src.MapKeys() {
//				newKey := reflect.New(key.Type()).Elem()
//				copyRecursive(newKey, key)
//				newValue := reflect.New(src.MapIndex(key).Type()).Elem()
//				copyRecursive(newValue, src.MapIndex(key))
//				dst.SetMapIndex(newKey, newValue)
//			}
//		}
//	default:
//		dst.Set(src)
//	}
//}
