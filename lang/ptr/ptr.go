package ptr

import "reflect"

// Of 将给定值转换为指针类型，如果输入的是空指针则返回nil
// 参数:
//
//	t: 任意类型的值
//
// 返回值:
//
//	*T: 指向输入值的指针，如果输入是空指针则返回nil
func Of[T any](t T) *T {
	rv := reflect.ValueOf(t)
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return nil
	}
	return &t
}

// From 将指针解引用为值类型，如果指针为nil则返回类型的零值
// 参数:
//
//	p: 指向任意类型值的指针
//
// 返回值:
//
//	T: 指针指向的值，如果指针为nil则返回类型的零值
func From[T any](p *T) T {
	if p != nil {
		return *p
	}
	var t T
	return t
}

func FromOrDefault[T any](p *T, def T) T {
	if p != nil {
		return *p
	}
	return def
}
