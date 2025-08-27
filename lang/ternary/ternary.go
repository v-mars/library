package ternary

// IFElse 示例：ternary.IFElse(score > 90, "true value", "false value")
// IFElse 是一个泛型三元操作函数，根据条件返回两个值中的一个
// 参数:
//   - ok: 布尔条件，决定返回哪个值
//   - trueValue: 当条件为true时返回的值
//   - falseValue: 当条件为false时返回的值
//
// 返回值:
//   - 如果ok为true，返回trueValue；否则返回falseValue
func IFElse[T any](ok bool, trueValue, falseValue T) T {
	if ok {
		return trueValue
	}
	return falseValue
}
