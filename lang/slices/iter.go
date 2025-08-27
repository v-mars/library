package slices

// Transform 将源切片中的每个元素通过转换函数映射为另一种类型的元素
// 参数:
//   - src: 源切片
//   - fn: 转换函数，将类型A转换为类型B
//
// 返回值:
//   - 转换后的新切片，类型为[]B
func Transform[A, B any](src []A, fn func(A) B) []B {
	if src == nil {
		return nil
	}

	dst := make([]B, 0, len(src))
	for _, a := range src {
		dst = append(dst, fn(a))
	}

	return dst
}

// TransformWithErrorCheck 将源切片中的每个元素通过可能返回错误的转换函数映射为另一种类型的元素
// 参数:
//   - src: 源切片
//   - fn: 转换函数，将类型A转换为类型B，可能返回错误
//
// 返回值:
//   - 转换后的新切片，类型为[]B
//   - 如果转换过程中发生错误，则返回错误信息
func TransformWithErrorCheck[A, B any](src []A, fn func(A) (B, error)) ([]B, error) {
	if src == nil {
		return nil, nil
	}

	dst := make([]B, 0, len(src))
	for _, a := range src {
		item, err := fn(a)
		if err != nil {
			return nil, err
		}
		dst = append(dst, item)
	}

	return dst, nil
}

// GroupBy 根据指定函数将切片元素分组
// 参数:
//   - src: 源切片
//   - fn: 分组函数，接收元素A，返回键K和值V
//
// 返回值:
//   - 分组后的映射，键为K类型，值为V类型的切片
//
// 示例: 按照年龄将人员分组
//
//	type Person struct {
//	   Name string
//	   Age  int
//	}
//
//	people := []Person{
//	   {"Alice", 25},
//	   {"Bob", 30},
//	   {"Charlie", 25},
//	   {"David", 30},
//	}
//
//	grouped := slices.GroupBy(people, func(p Person) (int, Person)) {
//	   return p.Age, p
//	})
//
// 结果: map[25:[{Alice 25} {Charlie 25}] 30:[{Bob 30} {David 30}]]
func GroupBy[A any, K comparable, V any](src []A, fn func(A) (K, V)) map[K][]V {
	if src == nil {
		return nil
	}
	dst := make(map[K][]V)
	for _, a := range src {
		k, v := fn(a)
		dst[k] = append(dst[k], v)
	}
	return dst
}

// Unique 去除切片中的重复元素，保留唯一元素
// 参数:
//   - src: 源切片
//
// 返回值:
//   - 去重后的新切片
func Unique[T comparable](src []T) []T {
	if src == nil {
		return nil
	}
	dst := make([]T, 0, len(src))
	m := make(map[T]struct{}, len(src))
	for _, s := range src {
		if _, ok := m[s]; ok {
			continue
		}
		dst = append(dst, s)
		m[s] = struct{}{}
	}

	return dst
}

// Fill 创建一个指定长度并用指定值填充的切片
// 参数:
//   - val: 用于填充的值
//   - size: 切片长度
//
// 返回值:
//   - 填充后的切片
//
// 示例：创建长度为5，用"hello"填充的字符串切片
// filled := slices.Fill("hello", 5)
// 结果: ["hello", "hello", "hello", "hello", "hello"]
func Fill[T any](val T, size int) []T {
	slice := make([]T, size)
	for i := 0; i < size; i++ {
		slice[i] = val
	}
	return slice
}

// Chunks 将切片分割成指定大小的块
// 参数:
//   - s: 源切片
//   - chunkSize: 每个块的大小
//
// 返回值:
//   - 分割后的二维切片
//
// 示例：将整数切片分割成大小为3的块
// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// chunks := slices.Chunks(numbers, 3)
// 结果: [[1, 2, 3], [4, 5, 6], [7, 8, 9], [10]]
func Chunks[T any](s []T, chunkSize int) [][]T {
	sliceLen := len(s)
	chunks := make([][]T, 0, sliceLen/chunkSize)

	for start := 0; start < sliceLen; start += chunkSize {
		end := start + chunkSize
		if end > sliceLen {
			end = sliceLen
		}

		chunks = append(chunks, s[start:end])
	}

	return chunks
}

// ToMap 将切片转换为映射
// 参数:
//   - src: 源切片
//   - fn: 转换函数，接收元素E，返回键K和值V
//
// 返回值:
//   - 转换后的映射
//
// 示例：将人员切片转换为以姓名为键的映射
//
//	people := []Person{
//	   {"Alice", 25},
//	   {"Bob", 30},
//	   {"Charlie", 35},
//	}
//
//	peopleMap := slices.ToMap(people, func(p Person) (string, Person) {
//	   return p.Name, p
//	})
//
// 结果: map[Alice:{Alice 25} Bob:{Bob 30} Charlie:{Charlie 35}]
func ToMap[E any, K comparable, V any](src []E, fn func(e E) (K, V)) map[K]V {
	if src == nil {
		return nil
	}

	dst := make(map[K]V)
	for _, e := range src {
		k, v := fn(e)
		dst[k] = v
	}

	return dst
}

// Reverse 反转切片中元素的顺序
// 参数:
//   - slice: 要反转的切片
//
// 返回值:
//   - 反转后的切片（原地反转）
//
// 示例：反转整数切片
// numbers := []int{1, 2, 3, 4, 5}
// reversed := slices.Reverse(numbers)
// // 结果: [5, 4, 3, 2, 1]
// // 注意：这是原地反转，numbers切片本身也会被修改
func Reverse[T any](slice []T) []T {
	left := 0
	right := len(slice) - 1
	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
	return slice
}
