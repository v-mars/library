package maps

// ToAnyValue 将键值对映射中的值类型转换为interface{}类型
// 参数:
//
//	m: 键为可比较类型、值为任意类型的映射
//
// 返回值:
//
//	map[K]any: 键为相同类型、值为interface{}类型的映射
func ToAnyValue[K comparable, V any](m map[K]V) map[K]any {
	// 创建新的映射，容量与原映射相同
	n := make(map[K]any, len(m))
	// 遍历原映射，将键值对复制到新映射中
	for k, v := range m {
		n[k] = v
	}

	return n
}

// TransformKey 将映射中的键类型进行转换
// 参数:
//
//	m: 原始映射，键类型为K1，值类型为V
//	f: 键转换函数，将K1类型转换为K2类型
//
// 返回值:
//
//	map[K2]V: 转换后的新映射，键类型为K2，值类型保持不变
func TransformKey[K1, K2 comparable, V any](m map[K1]V, f func(K1) K2) map[K2]V {
	// 创建新的映射，容量与原映射相同
	n := make(map[K2]V, len(m))
	// 遍历原映射，使用转换函数处理键并构建新映射
	for k1, v := range m {
		n[f(k1)] = v
	}
	return n
}
