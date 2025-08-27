package sets

// Set 是一个泛型集合类型，使用map实现，值为空结构体以节省内存
type Set[T comparable] map[T]struct{}

// FromSlice 从切片创建集合，去除重复元素
// 参数:
//
//	s: 输入的切片
//
// 返回值:
//
//	Set[T]: 包含切片中所有唯一元素的集合
func FromSlice[T comparable](s []T) Set[T] {
	set := make(Set[T], len(s))
	// 遍历切片，将每个元素添加到集合中
	for _, elem := range s {
		cpElem := elem
		set[cpElem] = struct{}{}
	}
	return set
}

// ToSlice 将集合转换为切片
// 返回值:
//
//	[]T: 包含集合中所有元素的切片
func (s Set[T]) ToSlice() []T {
	// 创建切片，容量为集合大小
	sl := make([]T, 0, len(s))
	// 遍历集合并将元素添加到切片中
	for elem := range s {
		cpElem := elem
		sl = append(sl, cpElem)
	}
	return sl
}

// Contains 检查集合是否包含指定元素
// 参数:
//
//	elem: 要检查的元素
//
// 返回值:
//
//	bool: 如果集合包含该元素返回true，否则返回false
func (s Set[T]) Contains(elem T) bool {
	_, ok := s[elem]
	return ok
}

// Add 向集合添加元素
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// Remove 从集合中移除元素
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// Size 返回集合大小
func (s Set[T]) Size() int {
	return len(s)
}

// IsEmpty 检查集合是否为空
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// Clear 清空集合
func (s Set[T]) Clear() {
	for elem := range s {
		delete(s, elem)
	}
}

// Union 并集操作
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T])
	for elem := range s {
		result[elem] = struct{}{}
	}
	for elem := range other {
		result[elem] = struct{}{}
	}
	return result
}

// Intersection 交集操作
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T])
	// 选择较小的集合作为遍历对象以提高效率
	smaller, larger := s, other
	if len(s) > len(other) {
		smaller, larger = other, s
	}
	for elem := range smaller {
		if larger.Contains(elem) {
			result[elem] = struct{}{}
		}
	}
	return result
}

// Difference 差集操作 (s - other)
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := make(Set[T])
	for elem := range s {
		if !other.Contains(elem) {
			result[elem] = struct{}{}
		}
	}
	return result
}
